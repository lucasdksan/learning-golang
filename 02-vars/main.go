package main

import "fmt"

func main() {
	var varString string = "Variável 1"
	varString2 := "Variável 2" // Variável implícita ou inferência de tipo

	var (
		varString3 string = "Variável 3"
		varString4 string = "Variável 4"
	)

	varString5, varString6 := "Variável 5", "Variável 6"

	fmt.Println(varString)
	fmt.Println(varString2)
	fmt.Println(varString3)
	fmt.Println(varString4)
	fmt.Println(varString5)
	fmt.Println(varString6)

	/*
		Existe o var e o const a diferença entre eles é a mesma no JavaScript
	*/

	/*
		Forma simples de trocar os valores entre variáveis no GOlang
	*/

	varString5, varString6 = varString6, varString5

	fmt.Println(varString5)
	fmt.Println(varString6)
}
