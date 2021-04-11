package aladin

type Options struct {
	source Source
	parser Parser
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	return Options{}
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
