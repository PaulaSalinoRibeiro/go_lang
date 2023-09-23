package main

import "fmt"

func soma(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1 int, n2 int) (int, int) {
	soma := n1 + n2
	subtr := n1 - n2

	return soma, subtr
}

func main() {
	soma := soma(2, 2)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto")

	sm, sb := calculos(10, 2)
	fmt.Println(sm, sb)
}
