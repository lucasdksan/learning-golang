package main

import "fmt"

var n int

func main() {
	fmt.Println("Função main executada", n)
}

func init() {
	fmt.Println("Função init executada")
	n = 190
}
