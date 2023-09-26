package main

import "fmt"

func calc(n1 int, n2 int) (soma int, sub int) {
	// não é necessário utilizar o operador de atribuição :=
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calc(2, 3)
	fmt.Println(soma, sub)
}
