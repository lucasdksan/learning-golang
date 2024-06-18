package main

import "fmt"

func set_number_in_channel(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i
	}

	close(channel)
}

func main() {
	channel := make(chan int)

	go set_number_in_channel(channel)

	for num := range channel {
		fmt.Println(num)
	}
}
