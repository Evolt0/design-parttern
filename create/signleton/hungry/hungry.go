package hungry

// Singleton 饿汉式
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func NewHungryInstance() *Singleton {
	return singleton
}

