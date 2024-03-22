package main

import "fmt"

func main() {
	soma := 10 + 22
	sub := 10 - 5
	div := 10 / 5
	mult := 10 * 2
	rest := 10 % 3

	fmt.Println(soma, sub, div, mult, rest)

	// Atribuição

	var variavel string = "Lucas"
	var_2 := "Lucas"

	fmt.Println(variavel, var_2)

	// Operadores

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)

	// Operadores Logicos

	ver, fal := true, false

	fmt.Println(ver && fal)
	fmt.Println(fal || ver)
	fmt.Println(!ver)

	// Operadores unários
	number := 10
	number++
	fmt.Println(number)
	number += 15
	fmt.Println(number)

	num_2 := 10
	num_2--
	fmt.Println(num_2)
	num_2 -= 15
	fmt.Println(num_2)
}
