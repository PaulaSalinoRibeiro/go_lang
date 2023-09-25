package main

import "fmt"

func main() {
	v1 := 10
	v2 := v1 // fez uma cópia

	fmt.Println(v1, v2)
	v1++
	fmt.Println(v1, v2) // output 11 10

	// ponteiro é uma referencia de memória
	v3 := 100
	var pont *int
	fmt.Println(v3, pont) // output 100 <nil>

	pont = &v3
	fmt.Println(v3, pont) // output 100 0x1400009c020

	fmt.Println(v3, *pont) // output 100 100 - desrefenciação

	v3 = 150
	fmt.Println(v3, *pont) // output 150 150
}
