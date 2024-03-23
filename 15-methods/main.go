package main

import "fmt"

type user struct {
	name string
	year uint
}

func (u user) save() {
	fmt.Printf("Salvando os dados do user %s no banco de dados", u.name)
}

func (u user) check_age_majority() bool {
	return u.year >= 18
}

func (u *user) plus_age() {
	u.year++
}

func main() {
	user_l := user{name: "Lucas da Silva", year: 26}
	user_l.save()
	println("Maior idade?", user_l.check_age_majority())
	user_l.plus_age()

	fmt.Println("Nova idade: ", user_l)
}
