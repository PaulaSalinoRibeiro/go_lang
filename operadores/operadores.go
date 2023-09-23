package main

import "fmt"

func main() {

	fmt.Println(2 + 2)
	fmt.Println(2 - 2)
	fmt.Println(2 * 2)
	fmt.Println(2 / 2)
	fmt.Println(2 % 2)

	fmt.Println(2 > 2)
	fmt.Println(2 >= 2)
	fmt.Println(2 < 2)
	fmt.Println(2 <= 2)
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
}
