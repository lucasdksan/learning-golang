package main

import (
	"fmt"
)

// Escreva uma função em Go que receba uma string como parâmetro e retorne o número de vogais nessa string.

func validation(text string) uint {
	text_len := len(text)

	return uint(text_len)
}

func main() {
	var text string

	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&text)
	result := validation(text)

	fmt.Print("Número de letras: ", result)
}
