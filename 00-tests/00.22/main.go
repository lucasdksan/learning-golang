package main

import (
	"fmt"
	"strconv"
)

func convert_celsiuc_fahrenheit(value int64) float64 {
	let_number := float64(9) / 5
	return let_number*float64(value) + 32
}

func main() {
	var number string

	fmt.Println("Digite a temperatura em Celsius: ")
	fmt.Scan(&number)

	value, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Erro ao converter entrada para número:", err)
		return
	}

	result := convert_celsiuc_fahrenheit(int64(value))

	fmt.Println("Resultado: ", result)
}
