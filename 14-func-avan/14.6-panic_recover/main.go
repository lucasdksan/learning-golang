package main

import "fmt"

func recover_flux() {
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada")
	}
}

func student_media(grades ...float64) bool {
	defer fmt.Println("Está calculando a média")
	defer recover_flux()
	result_sum := float64(0)
	grades_len := len(grades)

	for _, grade := range grades {
		result_sum += grade
	}

	result_grade := result_sum / float64(grades_len)

	if result_grade > 6 {
		return true
	} else if result_grade < 6 {
		return false
	}

	panic("A MEDIA É 6")
}

func main() {
	fmt.Println("Defer") // Adiar até o ultimo momento possivel

	result := student_media(6, 6)

	fmt.Println("Result: ", result)
}
