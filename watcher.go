package aladin

// Watcher 配置监听器，监听配置内容更新
type Watcher interface {
	// Next 获取新配置快照
	Next() (*Snapshot, error)
	// Stop 停止监听配置
	Stop() error
}
