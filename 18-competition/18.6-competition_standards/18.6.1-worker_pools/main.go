package main

import "fmt"

func main() {
	to_dos := make(chan int, 45)
	results := make(chan int, 45)

	go worker(to_dos, results)

	for i := 0; i < 45; i++ {
		to_dos <- i
	}

	close(to_dos)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(to_dos <-chan int, results chan<- int) {
	for number := range to_dos {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
