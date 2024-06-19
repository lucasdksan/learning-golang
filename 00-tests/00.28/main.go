package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int, 5)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			channel <- i
		}

		close(channel)
	}()

	go func() {
		for value_in_channel := range channel {
			fmt.Println("Valor: ", value_in_channel)
		}

		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("Todas as goroutines concluídas")
}
