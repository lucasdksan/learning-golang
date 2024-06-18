package main

import (
	"fmt"
	"time"
)

func send_string_channel_1(channel chan string) {
	var i = 0
	for {
		channel <- fmt.Sprintf("Mensagem do canal 1 %d", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

func send_string_channel_2(channel chan string) {
	var i = 0
	for {
		channel <- fmt.Sprintf("Mensagem do canal 2 %d", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go send_string_channel_1(channel_1)
	go send_string_channel_2(channel_2)

	for {
		select {
		case msg1 := <-channel_1:
			fmt.Println(msg1)
		case msg2 := <-channel_2:
			fmt.Println(msg2)
		}
	}
}
