package aladin

// Parser 解析Snapshot
// 每个方法都要求 concurrent-safe
type Parser interface {
	// Parse 将Snapshot转换为Store
	Parse(*Snapshot) (Store, error)
}

// Store 存储经Parser处理过的配置数据
type Store interface {
	// Get 获取配置分支值，如对于下面配置
	// {"foo":{"bar":123}}
	// 调用 Get("foo.bar").Int(0) 返回的值为123
	Get(string) Value
	// Set 设置配置分支值，如对于下面配置
	// {"foo":{"bar":123}}
	// 调用 Set("foo.bar", 456) 后，配置变为
	// {"foo":{"bar":456}}
	Set(string, interface{})
	// Del 删除配置分支，如对于下面配置
	// {"foo":{"bar":123}}
	// 调用 Del("foo.bar") 后，配置变为
	// {"foo":{"bar":null}}
	Del(string)
	// Scan 将Snapshot 反序列化到结构体或map
	Scan(interface{}) error
}
