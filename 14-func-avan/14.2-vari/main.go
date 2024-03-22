package main

import "fmt"

func sum(b ...int) int {
	total := 0

	for _, value := range b {
		total += value
	}

	return total
}

func write(text string, numbers ...int) { // A prop variatica deve ser sempre a ultima e só pode existir uma!
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println("var")
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Total: ", total)

	write("Olá: ", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
