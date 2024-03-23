package main

import "fmt"

func closure() func() {
	text := "Dentro da função closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Dentro da Main"
	fmt.Println(text)

	new_function := closure()
	new_function()
}
