package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 Tipos de números que o GO suporta. Eles variam conforme a quantidade de Bits.
	// Caso coloque apenas o int, o GO vai considerar a arquitetura do seu PC como base. Exemplo, caso seu computador seja 64bits ele vai se adaptar para int64 e por ai vai...

	const number int = 105005055050

	// Caso utilize o uint sua variavel não vai permitir números negativos.

	fmt.Println("Oi Lucas!", number)

	// Os números reais podem ser representados pelos tipos float32 ou float64

	numFloat := 561.44

	fmt.Println("Oi Lucas!", numFloat)

	// Strings

	const str string = "Oi lucas, tudo bem?"

	fmt.Println(str)

	str2 := "texto"

	fmt.Println(str2)

	char := 'B' // Usando aspas simples vai transformar sua variável no tipo number, pois vai pegar o valor numerico da letra na tabela ascii. Só pode utilizar uma letra nas aspas simples.

	fmt.Println(char)

	// Todo o tipo de dado tem seu valor inicial

	var (
		newNumber int
		newChar   string
	)

	fmt.Println(newNumber)
	fmt.Println(newChar)

	const boolexemplo bool = false

	var errorexemplo error
	var errorExemplos error = errors.New("Error Interno")

	fmt.Println(boolexemplo)
	fmt.Println(errorexemplo)
	fmt.Println(errorExemplos)
}
