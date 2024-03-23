package main

import "fmt"

func invert_signal(number int) int {
	return number * -1
}

func invert_signal_with_pointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 100
	invert_number := invert_signal(number)
	invert_signal_with_pointer(&number)
	fmt.Println("Número invertido: ", invert_number)
	fmt.Println("Número invertido via ponteiro: ", number)
}
