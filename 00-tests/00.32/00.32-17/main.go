package main

import "fmt"

func main() {
	var arr = [10]int{10, 9, 8, 70, 6, 5, 4, 3, 2, 1}
	var arr_five [5]int

	copy(arr_five[:], arr[:5])

	fmt.Println(arr_five)
}
