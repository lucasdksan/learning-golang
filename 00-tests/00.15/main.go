package main

import (
	"fmt"
	"math"
)

// Escreva uma função em Go que receba um número inteiro como parâmetro e retorne verdadeiro se o número for um quadrado perfeito e falso caso contrário.

func main() {
	var number int

	fmt.Println("Digite um numéro para determinar se é um quadrado perfeito: ")
	fmt.Scan(&number)

	result := math.Sqrt(float64(number))

	if result == float64(int(result)) {
		fmt.Println("É um quadrado perfeito")
	} else {
		fmt.Println("Não é um quadrado perfeito")
	}
}
