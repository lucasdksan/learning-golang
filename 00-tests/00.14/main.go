package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Função para concatenar uma lista de strings
func concat(slice []string) string {
	return strings.Join(slice, " ")
}

func main() {
	var input string
	strings_slice := []string{}

	fmt.Println("Digite uma string, ou digite 'x' para parar:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()

		if input == "x" {
			break
		} else {
			strings_slice = append(strings_slice, input)
		}
	}

	result := concat(strings_slice)

	fmt.Println("String completa:", result)
}
