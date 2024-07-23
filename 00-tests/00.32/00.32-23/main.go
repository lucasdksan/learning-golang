package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	pi := math.Pi
	result := pi * math.Pow(c.radius, 2)

	return math.Round(result)
}

func main() {
	cir := circle{radius: 12.2}

	fmt.Print(cir.area())
}
