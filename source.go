package aladin

// Source 配置源，包括
// 本地文件配置；
// 远程配置服务；
// 环境变量配置；
// Redis 存储的配置等
type Source interface {
	// Read 读取配置
	Read() (*Snapshot, error)
	// Write 写入配置
	Write(*Snapshot) error
	// Watch 获取配置监听器
	Watch() (Watcher, error)
	// Close 关闭配置源
	Close() error
	// String 返回配置源类型
	String() string
}

// Watcher 配置监听器，监听配置内容更新
type Watcher interface {
	// Next 获取新配置快照
	Next() (*Snapshot, error)
	// Stop 停止监听配置
	Stop() error
}
