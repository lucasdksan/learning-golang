package main

import "fmt"

// Escreva uma função em Go que receba uma string e retorne o número de palavras na string.

func count_string(text string) int {
	return len(text)
}

func main() {
	var text string

	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&text)

	result := count_string(text)

	fmt.Print(result)
}
