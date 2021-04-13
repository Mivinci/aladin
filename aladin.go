package aladin

var DefaultConfig Config

func Init(opts ...Option) (err error) {
	DefaultConfig, err = New(opts...)
	return
}

func Sync() error {
	return DefaultConfig.Sync()
}

func Snapshots() *Snapshot {
	return DefaultConfig.Snapshot()
}

func Close() error {
	return DefaultConfig.Close()
}

func Get(path string) Value {
	return DefaultConfig.Get(path)
}

func Set(path string, v interface{}) {
	DefaultConfig.Set(path, v)
}

func Del(path string) {
	DefaultConfig.Del(path)
}

func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}
