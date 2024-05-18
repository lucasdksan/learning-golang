package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Escreva uma função em Go que receba uma lista de inteiros e retorne a soma de todos os números na lista.

func main() {
	var input string
	var result uint64 = 0
	numbers_slice := []uint64{}

	fmt.Println("Digite uma string, ou digite 'x' para parar:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()

		if input == "x" {
			break
		} else {
			number, err := strconv.ParseUint(input, 10, 64)
			if err != nil {
				fmt.Println("Erro ao converter entrada para número:", err)
				continue
			}
			numbers_slice = append(numbers_slice, number)
		}
	}

	for _, number := range numbers_slice {
		result = result + number
	}

	fmt.Println("Soma de todos os números: ", result)
}
