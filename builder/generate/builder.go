package generate

const (
	defaultData = 20
)

type Builder struct {
	data int
}

type Config struct {
	data int
}

func (b *Builder) SetData(data int) {
	b.data = data
}

func (b *Builder) Build() *Config {
	if b.data == 0 {
		b.data = defaultData
	}
	return &Config{
		data: b.data,
	}
}
