package main

import "fmt"

func function_1() {
	fmt.Println("Função 1")
}

func function_2() {
	fmt.Println("Funçãio 2")
}

func student_media(grades ...float64) float64 {
	defer fmt.Println("Está calculando a média")
	result := float64(0)
	grades_len := len(grades)

	for _, grade := range grades {
		result += grade
	}

	return result / float64(grades_len)
}

func main() {
	fmt.Println("Defer") // Adiar até o ultimo momento possivel

	defer function_1()

	function_2()
	result := student_media(10, 15.5, 1.5)

	fmt.Println("Result: ", result)
}
