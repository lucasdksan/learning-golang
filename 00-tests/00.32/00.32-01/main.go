package main

func is_even(number int) bool {
	return number%2 == 0
}

func main() {
	result := is_even(13)

	println("Resultado: ", result)
}
