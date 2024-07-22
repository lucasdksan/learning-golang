package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (obj rectangle) area() int {
	return obj.height * obj.width
}

func (obj rectangle) perimeter() int {
	return 2*obj.height + 2*obj.width
}

func main() {
	rec := rectangle{width: 135, height: 23}

	fmt.Print(rec.perimeter())
}
