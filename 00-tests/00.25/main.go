package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	for i := 0; i < 3; i++ {
		channel <- i
	}

	close(channel)

	for num := range channel {
		fmt.Println(num)
	}
}
