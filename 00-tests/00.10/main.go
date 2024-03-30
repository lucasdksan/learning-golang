package main

import (
	"fmt"
)

// Escreva um programa em Go que leia um número inteiro do usuário e imprima todos os números primos menores ou iguais a esse número.
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func printPrimes(n int) {
	fmt.Println("Números primos menores ou iguais a", n, ":")
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	printPrimes(num)
}
