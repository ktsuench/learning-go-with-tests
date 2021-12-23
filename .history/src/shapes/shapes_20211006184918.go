package shapes

import "math"

type Shape interface {
	Area() floatt64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
