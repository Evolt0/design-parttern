package adapter

type IServer interface {
	Create(cpu, mem string)
}

type Aws struct {
	cpu string
	mem string
}

func NewAws(cpu string, mem string) *Aws {
	return &Aws{cpu: cpu, mem: mem}
}

type AwsAdapter struct {
	aws Aws
}

func (a *AwsAdapter) Create(cpu, mem string) {
	aws := NewAws(cpu, mem)
	a.aws = *aws
}

type Aliyun struct {
	cpu string
	mem string
}

func NewAliyun(cpu string, mem string) *Aliyun {
	return &Aliyun{cpu: cpu, mem: mem}
}

type AliyunAdapter struct {
	aliyun Aliyun
}

func (a *AliyunAdapter) Create(cpu, mem string) {
	aliyun := NewAliyun(cpu, mem)
	a.aliyun = *aliyun
}
