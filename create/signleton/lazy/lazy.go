package lazy

import "sync"

type Singleton struct{}

var (
	singleton *Singleton
	once      = &sync.Once{}
)

// NewLazyInstance 懒汉式
func NewLazyInstance() *Singleton {
	if singleton == nil {
		once.Do(func() {
			singleton = &Singleton{}
		})
	}
	return singleton
}
