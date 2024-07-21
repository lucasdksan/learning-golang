package main

import "fmt"

func get_strong_number(arr [10]int) int {
	var result int = 0

	for i := 0; i < len(arr); i++ {
		if result < arr[i] {
			result = arr[i]
		}
	}

	return result
}

func main() {
	var arr = [10]int{10, 9, 8, 70, 6, 5, 4, 3, 2, 1}

	result := get_strong_number(arr)

	fmt.Println(result)
}
