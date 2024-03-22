package main

import "fmt"

func main() {
	fmt.Println("var")
	func(text string) {
		fmt.Println(text)
	}("Olá Jovens")

	funExec := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Olá Lucas")

	fmt.Println(funExec)

	funExecSum := func(a int, b int) int {
		return a + b
	}

	result := funExecSum(1, 3)

	fmt.Println("Soma: ", result)
}
