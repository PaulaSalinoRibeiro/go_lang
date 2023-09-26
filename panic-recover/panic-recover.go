package main

import "fmt"

func recupera() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func soma(n ...int) int {
	defer recupera()
	total := 0
	for _, v := range n {
		total += v
	}
	if total > 10 {
		panic("Total é maior que 10")
	}
	return total
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5))
	fmt.Println("Fim da execusão do programa")
}
