package main

func factorial(number int) int {
	result := number

	for i := number - 1; i > 0; i-- {
		println("I: ", i)
		result = result * i
	}

	return result
}

func main() {
	result := factorial(5)

	print(result)
}
