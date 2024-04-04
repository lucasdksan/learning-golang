package main

import (
	"fmt"
	"strconv"
)

// Escreva uma função em Go que receba uma lista de inteiros e retorne uma nova lista contendo apenas os números pares da lista original.

func create_new_arr(arr []int) []int {
	var new_arr []int

	for _, number := range arr {
		if number%2 == 0 {
			new_arr = append(new_arr, number)
		}
	}

	return new_arr
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

	result := create_new_arr(slice_numbers)

	fmt.Println("O nova array: ", result)
}
