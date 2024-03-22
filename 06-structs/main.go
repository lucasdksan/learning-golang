package main

import "fmt"

type user struct {
	name    string
	year    uint8
	address address
}

type address struct {
	city         string
	uf           string
	street       string
	number       int
	neighborhood string
	cep          string
}

func main() {
	var us user
	addr := address{"Parnamirim", "RN", "Rua Capitão Martinho Machado", 1859, "Passagem de Áreia", "40028922"}

	us.name = "Lucas da Silva Leoncio"
	us.year = 26
	us.address = addr

	us2 := user{"Lucas Silva", 26, addr}
	us3 := user{name: "Lucas"}

	fmt.Println("Start Struct", us3)
	fmt.Println("Start Struct", us2)
	fmt.Println("Start Struct", us)
}
