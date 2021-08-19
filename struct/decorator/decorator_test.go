package decorator

import (
	"testing"
)

func TestNewColorSquare(t *testing.T) {
	square := Square{}
	colorSquare := NewColorSquare(square, "yellow")
	t.Log(colorSquare.Draw())
}
