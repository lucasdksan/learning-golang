package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexar(write("Lucas"), write("Silva"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexar(channel_input_1, channel_input_2 <-chan string) <-chan string {
	channel_output := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel_input_1:
				channel_output <- message
			case message := <-channel_input_2:
				channel_output <- message
			}
		}
	}()

	return channel_output
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
