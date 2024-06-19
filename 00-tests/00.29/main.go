package main

import (
	"fmt"
	"time"
)

func send_int_in_channel(channel chan int) {
	time.Sleep(1 * time.Second)
	channel <- 42
}

func main() {
	channel := make(chan int)

	go send_int_in_channel(channel)

	select {
	case num := <-channel:
		fmt.Println("Recebido:", num)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout expirado antes de receber")
	}
}
