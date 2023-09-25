package main

import "fmt"

func main() {
	n := 5

	if n > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor ou igual a 5")
	}

	if num := n; num > 0 {
		fmt.Println("Numero maior que zero")
	} else {
		fmt.Println("Numero menor ou igual a zero")
	}
}
