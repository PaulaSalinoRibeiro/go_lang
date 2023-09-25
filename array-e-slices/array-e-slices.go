package main

import (
	"fmt"
)

func main() {
	// array
	var arr1 [5]int
	fmt.Println(arr1) // output [0 0 0 0 0] preenche com valor zero
	// add items
	arr1[0] = 1
	fmt.Println(arr1) // output [1 0 0 0 0]
	// inicializa com os valores
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2) // output [1 2 3 4 5]

	// slice
	slc := []int{10, 11, 12, 13, 14}
	fmt.Println(slc) // output [10 11 12 13 14]

	// add items
	slc = append(slc, 15)
	fmt.Println(slc) // output [10 11 12 13 14 15]

	slc2 := arr2[1:3] // indice 1 inclusivo e indice 3 exclusivo e passado como referencia
	fmt.Println(slc2) // output [2 3]

	arr2[1] = 20
	fmt.Println(slc2) //output [20 3]

	// array interno
	slc3 := make([]int, 10, 11)
	fmt.Println(slc3)      // output [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(slc3)) //output 10
	fmt.Println(cap(slc3)) // output 11

	slc3 = append(slc3, 12)
	slc3 = append(slc3, 13)

	fmt.Println(slc3)      // output [0 0 0 0 0 0 0 0 0 0 12 13]
	fmt.Println(len(slc3)) //output 12
	fmt.Println(cap(slc3)) // output 22
}
