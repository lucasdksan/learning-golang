package main

import "fmt"

type user struct {
	name     string
	email    string
	year     uint
	password string
}

type student struct {
	user
	course     string
	department string
}

func main() {
	fmt.Println("Oi Lucas")
}
