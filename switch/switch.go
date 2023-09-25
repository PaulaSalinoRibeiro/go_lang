package main

import "fmt"

func weekOfDay(n int) string {
	switch n {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func weekOfDay2(n int) string {
	switch {
	case n == 1:
		return "Domingo"
	case n == 2:
		return "Segunda"
	case n == 3:
		return "Terça"
	case n == 4:
		return "Quarta"
	case n == 5:
		return "Quinta"
	case n == 6:
		return "Sexta"
	case n == 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func main() {
	n := 3
	day := weekOfDay(n)
	fmt.Println(day)

	day2 := weekOfDay2(n)
	fmt.Println(day2)
}
