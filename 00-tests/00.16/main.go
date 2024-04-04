package main

import "fmt"

// Escreva um programa em Go que imprima os primeiros N termos da sequência de Fibonacci, onde N é fornecido pelo usuário.

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	var number int

	fmt.Println("Digite um numéro para calcular o Fibonacci: ")
	fmt.Scan(&number)

	result := fibonacci(uint(number))

	fmt.Println("Resultado final: ", result)
}
