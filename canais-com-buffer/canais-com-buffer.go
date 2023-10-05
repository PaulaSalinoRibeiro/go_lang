package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go"

	mensage := <-canal
	mensage2 := <-canal

	fmt.Println(mensage)
	fmt.Println(mensage2)

}
