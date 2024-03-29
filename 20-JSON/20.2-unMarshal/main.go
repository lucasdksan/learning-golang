package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	var c dog
	c_json := `{"name":"Eren", "breed":"vira-lata", "age": 2}`

	if erro := json.Unmarshal([]byte(c_json), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
}
