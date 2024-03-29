package main

import (
	"fmt"
	"strconv"
)

// Escreva uma função em Go que receba uma lista de inteiros e retorne a média desses números.

func validation(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}

	total := slice[0]

	for _, number := range slice[1:] {
		total += number
	}

	return float64(total) / float64(len(slice))
}

func main() {
	var value string

	slice_numbers := []int{}

	fmt.Println("Digite um número, caso queira parar digite x: ")

	for {
		fmt.Scan(&value)

		if value == "x" || value == "X" {
			break
		} else {
			number, err := strconv.Atoi(value)

			if err != nil {
				fmt.Println("Erro ao converter número:", err)
				continue
			}

			slice_numbers = append(slice_numbers, number)
		}
	}

	result := validation(slice_numbers)

	fmt.Println("A média desse array é: ", result)
}
