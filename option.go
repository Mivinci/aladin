package aladin

type Options struct {
	source    Source
	parser    Parser
	readOnly  bool
	hotReload bool
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	options := Options{
		source: NewFileSource(DefaultPath),
		parser: NewYAMLParser(),
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithSource(s Source) Option {
	return func(o *Options) {
		o.source = s
	}
}

func WithParser(p Parser) Option {
	return func(o *Options) {
		o.parser = p
	}
}

// WithReadOnly 使用该函数让config只可读，
// 即不能修改加载到内存中的配置值
func WithReadOnly() Option {
	return func(o *Options) {
		o.readOnly = true
	}
}

// WithHotReload 对于本地文件配置，使用该函数启动热更新，
// 即服务运行期手动更新配置文件内容会触发config重新读取配置
func WithHotReload() Option {
	return func(o *Options) {
		o.hotReload = true
	}
}

// WithEnvSource 使用环境变量配置源，等价于
// Init(
//     WithSource(NewEnvSource())
//     WithParser(NewEnvParser())
// )
func WithEnvSource() Option {
	return func(o *Options) {
		o.source = NewEnvSource()
		o.parser = NewEnvParser()
	}
}
