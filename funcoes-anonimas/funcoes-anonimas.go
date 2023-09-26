package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Parametro da função anonima")

}
