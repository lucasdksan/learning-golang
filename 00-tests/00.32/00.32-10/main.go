package main

import "fmt"

func main() {
	city := map[string]int{
		"São Paulo": 11,
		"Ceará":     88,
	}

	city["Rio Grande do Norte"] = 10

	fmt.Println(city)
}
