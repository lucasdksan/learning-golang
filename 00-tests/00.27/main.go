package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Mensagem da goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Mensagem da goroutine 2")
	}()

	wg.Wait()
	fmt.Println("Todas as goroutines concluídas")
}
