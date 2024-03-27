package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)
	channel <- "Olá Mundo"
	channel <- "Olá Mundo 2"

	message := <-channel
	message_2 := <-channel

	fmt.Printf(message)
	fmt.Printf(message_2)
}
