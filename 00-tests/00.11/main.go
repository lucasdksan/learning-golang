package main

import (
	"fmt"
	"strings"
)

// A função verifica se a string fornecida é um palíndromo.
func isPalindrome(text string) bool {
	// Normalizar a string removendo espaços em branco e convertendo para minúsculas.
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))
	// Iterar pela metade da string para verificar se os caracteres correspondentes são iguais.
	for i := 0; i < len(text)/2; i++ {
		// Verificar se os caracteres correspondentes não são iguais.
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var text string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&text)

	if isPalindrome(text) {
		fmt.Println("É um palíndromo")
	} else {
		fmt.Println("Não é um palíndromo")
	}
}
