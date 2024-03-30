package main

import (
	"fmt"
)

// Escreva um programa em Go que solicite ao usuário seu nome e, em seguida, imprima "Olá, [nome]!" na tela.

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSpace(input)

	// fmt.Printf("Olá, %s!", input)

	var input string

	fmt.Print("Digite o seu nome: ")
	fmt.Scanf("%s", &input)

	fmt.Printf("Olá, %s!", input)
}
