package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {

	posicao := fibonacci(8)
	fmt.Println(posicao)
}
