package main

import "fmt"

type user struct {
	nome  string
	idade int
}

func (u user) saldacao() {
	fmt.Println("Olá meu nome é", u.nome)
}

func (u *user) fazAniversario() {
	u.idade++
}

func main() {
	u := user{nome: "Paula", idade: 33}
	u.saldacao()
	fmt.Println(u.idade)
	u.fazAniversario()
	fmt.Println(u.idade)
}
