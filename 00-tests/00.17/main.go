package main

import (
	"fmt"
)

// Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida.

func invert_string(text string) string {
	txt := []rune(text)
	ret := []rune{}

	for i := len(txt) - 1; i >= 0; i-- {
		ret = append(ret, txt[i])
	}

	return string(ret)
}

func main() {
	var text string

	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&text)

	result := invert_string(text)

	fmt.Print(result)
}
