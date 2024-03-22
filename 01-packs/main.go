package main

import (
	"fmt"
	"module/auxiliar"
	"module/sum"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no main.go")
	auxiliar.Escrevendo()
	erro := checkmail.ValidateFormat("123")

	sumResult := sum.SumDo(1, 2)

	fmt.Println("Res: ", sumResult)

	fmt.Println(erro)
}
