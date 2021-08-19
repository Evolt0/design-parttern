package lazy_test

import (
	"github.com/Evolt0/design_pattern/signleton/lazy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLazyInstance(t *testing.T) {
	assert.Equal(t, lazy.NewLazyInstance(), lazy.NewLazyInstance())
}

func BenchmarkNewLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if lazy.NewLazyInstance() != lazy.NewLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
