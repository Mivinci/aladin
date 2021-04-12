package aladin

import (
	"os"
	"strconv"
	"strings"
)

var (
	_ Source  = (*env)(nil)
	_ Watcher = (*envWatcher)(nil)
)

type env struct {
	next chan struct{}
}

func NewEnvSource() Source {
	var c env
	c.next = make(chan struct{})
	return &c
}

func (s *env) Read() (*Snapshot, error) {
	return nil, nil
}

// 通过向next管道发送消息，触发watcher.Next停止阻塞
// 进而使config获取新配置数据，若不阻塞该管道，config
// 会一直读，这样就傻逼了（原因见config.go第74行）。
func (s *env) Write(snap *Snapshot) error {
	s.next <- struct{}{}
	return nil
}

func (s *env) Watch() (Watcher, error) {
	return newEnvWatcher(s)
}

func (s *env) Close() error {
	close(s.next)
	return nil
}

func (s *env) String() string {
	return "env"
}

type envWatcher struct {
	s *env
}

func newEnvWatcher(s *env) (Watcher, error) {
	return &envWatcher{s: s}, nil
}

func (w *envWatcher) Next() (*Snapshot, error) {
	<-w.s.next // 阻塞，不然config会一直读
	return w.s.Read()
}

func (envWatcher) Stop() error { return nil }

type envParser struct{}

func NewEnvParser() Parser {
	return new(envParser)
}

func (p *envParser) Parse(snap *Snapshot) (Store, error) {
	return p, nil
}

func (p *envParser) Get(path string) Value {
	v := os.Getenv(formatKey(path))
	return stringValue(v)
}

func (p *envParser) Set(path string, v interface{}) {
	var s string
	switch t := v.(type) {
	case string:
		s = t
	case int:
		s = strconv.Itoa(t)
	default:
	}
	os.Setenv(formatKey(path), s)
}

func (p *envParser) Del(path string) {
	os.Unsetenv(formatKey(path))
}

func (p *envParser) Scan(v interface{}) error {
	return nil
}

func formatKey(k string) string {
	k = strings.ToUpper(k)
	return strings.ReplaceAll(k, ".", "_")
}
