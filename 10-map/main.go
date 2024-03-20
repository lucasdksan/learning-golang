package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Lucas da Silva",
		"year": "20",
	}

	user2 := map[string]map[string]string{
		"name": {
			"first": "Lucas",
			"last":  "Silva",
		},
		"course": {
			"name":       "Engenharia Mecatrônica",
			"university": "UFRN",
		},
	}

	fmt.Println(user)
	fmt.Println(user2)

	delete(user2, "name")

	fmt.Println(user2)

	user2["user"] = map[string]string{
		"name": "Lucas da Silva",
		"year": "26",
	}

	fmt.Println(user2)
}
