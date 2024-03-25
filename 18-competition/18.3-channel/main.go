package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Olá mundo!", channel)

	fmt.Println("Start Loop")

	// for { Forma menos elegante
	// 	message, open := <-channel

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range channel { // Forma mais elegante
		fmt.Println(message)
	}

	fmt.Println("End Loop")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
