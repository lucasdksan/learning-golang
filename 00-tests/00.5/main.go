package main

import (
	"fmt"
)

// Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for primo e falso caso contrário.

func validation(number int) string {
	var result int = 0

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			result++
		}
	}

	if result == 0 {
		return "O numéro digitado é primo"
	} else {
		return "O número digitado não é primo"
	}
}

func main() {
	var number int

	fmt.Scan(&number)
	result := validation(number)

	fmt.Print(result)
}
