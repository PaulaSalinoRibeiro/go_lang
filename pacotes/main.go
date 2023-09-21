package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo pacotes")
	auxiliar.Escrevendo()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
