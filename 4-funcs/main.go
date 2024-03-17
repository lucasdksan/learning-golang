package main

import (
	"fmt"
	"module/adds"
	mathematicalcalculations "module/mathematicalCalculations"
)

func main() {
	result := adds.Adds(1, 5)
	result_2, result_3 := mathematicalcalculations.MathematicalCalculations(10, 5) // Se colocar _ ele vai ocultar/ignorar o segundo retorno. Utilizado quando vc tem uma função que retorna mais de dois valores, porém vc deseja apenas um deles.

	fmt.Println("Resultado: ", result)
	fmt.Println("Resultado2: ", result_2)
	fmt.Println("Resultado3: ", result_3)
}
