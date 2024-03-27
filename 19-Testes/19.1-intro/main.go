package main

import (
	"fmt"
	"intro/address"
)

func main() {
	type_address := address.Address_type("Rua Capitão Martinho Machado")

	fmt.Println(type_address)
}
