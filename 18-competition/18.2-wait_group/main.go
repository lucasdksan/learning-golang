package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait_group sync.WaitGroup

	wait_group.Add(2)

	go func() {
		write("Olá mundo!")

		wait_group.Done()
	}()

	go func() {
		write("Olá Lucas")

		wait_group.Done()
	}()

	wait_group.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
