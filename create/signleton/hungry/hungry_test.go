package hungry_test

import (
	"github.com/Evolt0/design_pattern/signleton/hungry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHungryInstanceInstance(t *testing.T) {
	assert.Equal(t, hungry.NewHungryInstance(), hungry.NewHungryInstance())
}

func BenchmarkNewHungryInstanceInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if hungry.NewHungryInstance() != hungry.NewHungryInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
