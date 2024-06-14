package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var number string
	var arr []int

	fmt.Println("Digite um número: ")
	fmt.Scan(&number)

	value, err := strconv.Atoi(number)

	if err != nil {
		fmt.Println("Erro ao converter entrada para número:", err)
		return
	}

	limit := math.Sqrt(float64(value))
	convert_limit := int(limit)

	for i := 1; i <= convert_limit; i++ {
		if value%i == 0 {
			arr = append(arr, i)
			if i != value/i {
				arr = append(arr, value/i)
			}
		}
	}

	fmt.Println("Divisores:", arr)
}
