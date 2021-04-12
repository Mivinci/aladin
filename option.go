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

func WithReadOnly() Option {
	return func(o *Options) {
		o.readOnly = true
	}
}

func WithHotReload() Option {
	return func(o *Options) {
		o.hotReload = true
	}
}

func WithEnv() Option {
	return func(o *Options) {
		o.source = NewEnvSource()
		o.parser = NewEnvParser()
	}
}
