package main

import "fmt"

func main() {
	fmt.Print("Estrutura de Controle")

	number := 10

	if (number >= 5) && (number <= 15) {
		fmt.Println("O número é maior que 5")
	} else {
		fmt.Println("O número é menor que 5")
	}

	if new_number := number + 1; new_number > 10 {
		fmt.Println("O número é maior que 10", new_number)
	} else {
		fmt.Println("O número é menor que 10", new_number)
	}
}
