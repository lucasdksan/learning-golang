package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Olá mundo!")
	write("Olá Lucas")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
