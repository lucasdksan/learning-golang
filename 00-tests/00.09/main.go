package main

import (
	"fmt"
	"strings"
)

// Escreva uma função em Go que receba uma lista de strings e retorne uma nova lista com todas as strings em letras maiúsculas.

func validation(text string) string {
	return strings.ToLower(text)
}

func main() {
	var text string

	fmt.Scan(&text)

	result := validation(text)

	fmt.Println("Seu input em LowerCase: ", result)
}
