package main

import "fmt"

// Escreva uma função em Go que calcule e retorne o fatorial de um número inteiro fornecido como parâmetro.

func factorial(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}

	return result
}

func main() {
	var number int

	fmt.Println("Digite um numéro para calcular a fatorial: ")
	fmt.Scan(&number)

	result := factorial(number)

	fmt.Println("O resultado é: ", result)
}
