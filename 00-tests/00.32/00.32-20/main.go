package main

import "fmt"

func reverse_string_arr(arr []string) []string {
	var new_string []string

	for _, str := range arr {
		runes := []rune(str)

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		new_string = append(new_string, string(runes))
	}

	return new_string
}

func main() {
	var string_arr = []string{"Lucas", "Aline", "Leonardo"}

	result := reverse_string_arr(string_arr)

	fmt.Print(result)
}
