package main

import (
	"fmt"
	"sync"
)

func create_goroutine(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d executada\n", value)
}

func main() {
	var wg sync.WaitGroup
	goroutine_number := 5

	for i := 1; i <= goroutine_number; i++ {
		wg.Add(1)

		go create_goroutine(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todas as goroutines concluídas")
}
