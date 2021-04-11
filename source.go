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
}
