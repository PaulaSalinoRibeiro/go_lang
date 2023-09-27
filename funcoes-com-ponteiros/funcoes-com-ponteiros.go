package main

import "fmt"

func inverteSinal(n int) int {
	return n * -1
}

func inverteSinalComPonteiro(n *int) {
	*n = *n * -1
}

func main() {
	n1 := 10
	fmt.Println(inverteSinal(n1)) // passado como copia output -10
	fmt.Println(n1)               // output 10

	n2 := 20
	fmt.Println(n2)              //output 20
	inverteSinalComPonteiro(&n2) // passado como referencia
	fmt.Println(n2)              //output -20
}
