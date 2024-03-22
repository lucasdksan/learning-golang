package main

import "fmt"

func main() {
	// Arrays
	var arr1 [5]int

	arr1[0] = 10

	fmt.Println(arr1)

	arr2 := [3]string{"P1", "P2", "P3"}

	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr3)

	// Slices

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slice)

	slice = append(slice, 10)

	fmt.Println(slice)

	slice2 := arr3[0:3]

	fmt.Println(slice2)

	// Arrays Internos

	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) //Capacidade
}
