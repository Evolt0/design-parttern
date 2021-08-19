package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	subject := Subject{}
	observer := &Observer{data: "类型1"}
	subject.Register(observer)
	subject.Register(&Observer{data: "类型2"})
	subject.Notify("hello world!")
	subject.Remove(observer)
	subject.Notify("hello test")
}
