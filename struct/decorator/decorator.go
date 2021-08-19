package decorator

type IDraw interface {
	Draw() string
}

type Square struct{}

func (s *Square) Draw() string {
	return "this is square"
}

type ColorSquare struct {
	square Square
	color  string
}

func NewColorSquare(square Square, color string) *ColorSquare {
	return &ColorSquare{square: square, color: color}
}

func (c *ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
