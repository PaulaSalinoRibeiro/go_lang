package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func main() {
	var u usuario
	u.nome = "Teste"
	u.idade = 18

	fmt.Println(u)

	u2 := usuario{"Nome", 10}
	fmt.Println(u2)

	u3 := usuario{idade: 21}
	fmt.Println(u3)
}
