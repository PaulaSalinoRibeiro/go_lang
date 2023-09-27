package main

import "fmt"

func closure() func() {
	text := "texto da funcao closure"

	return func() {
		fmt.Println(text)
	}
}

func main() {
	cl := closure()
	cl()
}
