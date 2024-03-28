package forms

import (
	"math"
)

type Form interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
