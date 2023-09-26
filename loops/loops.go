package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	slice := []int{1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}
