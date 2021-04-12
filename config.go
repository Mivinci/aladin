package aladin

import (
	"log"
	"sync"
	"time"
)

// Config 配置引擎
type Config interface {
	Store
	// Init 初始化配置
	// 包括Snapshot初始化和Parser初始化
	Init(...Option) error
	// Sync 同步配置
	// 包括同步Snapshot。如必要，也同步Parser
	Sync() error
	// Close 关闭配置
	// 对于本地文件配置，则关闭文件；
	// 对于远程配置服务，则关闭连接；
	// 对于环境变量配置，不做操作。
	Close() error
}

var _ Config = (*config)(nil)

const maxRetry = 5

type config struct {
	options Options

	sync.RWMutex
	snap  *Snapshot
	store Store // 当指针用

	exit chan struct{}
}

func New(opts ...Option) (Config, error) {
	var c config

	c.Init(opts...)
	if c.options.hotReload {
		go c.deamon()
	}

	return &c, nil
}

func (c *config) Init(opts ...Option) error {
	c.options = newOptions(opts...)
	c.exit = make(chan struct{})
	return c.Sync()
}

func (c *config) Sync() error {
	snap, err := c.options.source.Read()
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	err = c.sync(snap)
	if err != nil {
		return err
	}
	return nil
}

func (c *config) sync(snap *Snapshot) (err error) {
	c.snap = snap
	c.store, err = c.options.parser.Parse(snap)
	return
}

func (c *config) deamon() {
	watch := func(w Watcher) error {
		for {
			snap, err := w.Next()
			if err != nil {
				return err
			}
			c.Lock()
			c.sync(snap)
			c.Unlock()
		}
	}
	for i := 0; i < maxRetry; i++ {
		w, err := c.options.source.Watch()
		if err != nil {
			log.Println("open watcher failed, retrying", i)
			time.Sleep(time.Second) // like it's retrying
			continue
		}

		done := make(chan struct{})

		go func() {
			select {
			case <-done:
			case <-c.exit:
			}
			w.Stop() // bp1
		}()

		if err = watch(w); err != nil {
			log.Panicln("next snapshot error", err)
		}

		close(done) // 触发bp1

		select {
		case <-c.exit: // 检查condig是否退出
			return
		default:
		}
	}
}

func (c *config) Close() error {
	close(c.exit)
	return c.options.source.Close()
}

func (c *config) Get(path string) Value {
	c.RLock()
	defer c.RUnlock()
	if c.store != nil {
		return c.store.Get(path)
	}
	return defValue()
}

func (c *config) Set(path string, v interface{}) {
	if c.options.readOnly {
		return
	}
	c.Lock()
	defer c.Unlock()
	if c.store != nil {
		c.store.Set(path, v)
	}
}

func (c *config) Del(path string) {
	if c.options.readOnly {
		return
	}
	c.Lock()
	defer c.Unlock()
	if c.store != nil {
		c.store.Del(path)
	}
}

func (c *config) Scan(v interface{}) error {
	c.RLock()
	defer c.RUnlock()
	if c.store != nil {
		return c.store.Scan(v)
	}
	return nil
}
