package main

import "fmt"

func main() {
	var val1 int = 10
	var val2 int = val1

	fmt.Println(val1, val2)

	val1++

	fmt.Println(val1, val2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var val3 int
	var p *int

	val3 = 3
	p = &val3

	fmt.Println(val3, p)
	fmt.Println(val3, *p) // Desreferenciação

	val3 += 15

	fmt.Println(val3, *p)
	fmt.Println(val3, p)
}
