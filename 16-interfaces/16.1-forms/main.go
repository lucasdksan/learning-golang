package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	heigth float64
	width  float64
}

func print_area(f form) {
	fmt.Printf("Á área da forma é %0.2f", f.area())
}

func (r rectangle) area() float64 {
	return r.heigth * r.width
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	r := rectangle{10, 20}
	c := circle{20}
	print_area(r)
	print_area(c)
}
