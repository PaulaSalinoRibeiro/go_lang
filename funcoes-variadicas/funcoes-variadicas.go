package main

import "fmt"

func soma(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}

	return total
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)
}
