package main

import (
	"bytes"
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
	c := dog{"eren", "vira-lata", 1}
	dog, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Print(dog)
	fmt.Print(bytes.NewBuffer(dog))
}
