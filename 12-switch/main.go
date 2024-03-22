package main

import "fmt"

func day_week(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return ""
	}
}

func day_week_2(day int) string {
	switch {
	case day == 1:
		return "Domingo"
		//  fallthrough Força o codigo ir para o proximo caso no switch
	case day == 2:
		return "Segunda-feira"
	case day == 3:
		return "Terça-feira"
	case day == 4:
		return "Quarta-feira"
	case day == 5:
		return "Quinta-feira"
	case day == 6:
		return "Sexta-feira"
	case day == 7:
		return "Sábado"
	default:
		return ""
	}
}

func main() {
	day := day_week(2)

	fmt.Println("Dia da Semana: ", day)
}
