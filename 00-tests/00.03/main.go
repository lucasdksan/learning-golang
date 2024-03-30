package main

import (
	"fmt"
)

// Escreva uma função em Go que receba dois números como parâmetros e retorne a soma desses números.

func main() {
	var number_1, number_2 int64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number_1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number_2)

	fmt.Printf("Resultado da soma: %d", number_1+number_2)
}
