package main

import "fmt"

func MathematicalCalculations(a, b int) (sum int, sub int) { // Uma vez declarando o que retornar não precisa criar as variaveis
	sum = a + b
	sub = a - b

	return // Alem disso não recisa deixar explicito o retorno
}

func main() {
	fmt.Println("Retorno Nomeado")

}
