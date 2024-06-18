package main

import "fmt"

func producer(channel chan int) {
	for i := 1; i < 11; i++ {
		channel <- i
	}

	close(channel)
}

func consumer(channel chan int, channel_done chan bool) {
	for len := range channel {
		fmt.Println(len)
	}

	channel_done <- true
}

func main() {
	channel := make(chan int)
	done := make(chan bool)

	go producer(channel)
	go consumer(channel, done)

	<-done
}
