package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		// time.Sleep(time.Second)
		fmt.Println("I: ", i)
	}

	for j := 0; j < 10; j++ {
		// time.Sleep(time.Second)
		fmt.Println("J: ", j)
	}

	names := [3]string{"Lucas", "Silva", "Leoncio"}

	for index, name := range names {
		fmt.Println("IndeX: ", index)
		fmt.Println("Name: ", name)
	}

	for index, letra := range "LUCAS" {
		fmt.Println("IndeX: ", index)
		fmt.Println("Name: ", string(letra))
	}

	user := map[string]string{
		"name":      "Lucas",
		"last_name": "Silva",
	}

	for index, props := range user {
		fmt.Println("IndeX: ", index)
		fmt.Println("Name: ", props)
	}

	// Não pode usar o range em structs

	// for { Loop infinito kkk!

	// }
}
