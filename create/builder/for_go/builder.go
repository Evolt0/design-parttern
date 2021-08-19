package for_go

type Builder struct {
	data int
}

type Option func(*Builder)

func WithData(data int) Option {
	return func(builder *Builder) {
		builder.data = data
	}
}

func NewBuilder(opts ...Option) *Builder {
	builder := &Builder{}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
}
