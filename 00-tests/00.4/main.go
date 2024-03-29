package main

import (
	"fmt"
)

// Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for par e falso caso contrário.

func validation(number int) string {
	if number%2 == 0 {
		return "O numéro digitado é par"
	} else {
		return "O número digitado é ímpar"
	}
}

func main() {
	var number int

	fmt.Scan(&number)
	result := validation(number)

	fmt.Print(result)
}
