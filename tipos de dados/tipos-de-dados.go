package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = 100000000
	fmt.Println(number)

	var numero2 uint = 10
	fmt.Println(numero2)

	// alias
	// rune = int32
	var numero3 rune = 123
	fmt.Println(numero3)

	// byte = unint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var str1 string = "Um texto"
	fmt.Println(str1)

	str2 := "Texto com inferencia de tipos"
	fmt.Println(str2)

	char := 'P'
	fmt.Println(char) // output 80 -> valor na tabela ascii

	boleano := true
	fmt.Println(boleano)

	var erro error = errors.New("novo erro")
	fmt.Println(erro)
}
