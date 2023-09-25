package main

import "fmt"

func main() {
	u := map[string]string{
		"nome":      "Paula",
		"sobrenome": "Ribeiro",
	}

	fmt.Println(u)
	fmt.Println(u["nome"])

	delete(u, "sobrenome")
	fmt.Println(u)

	u["idade"] = "33"
	fmt.Println(u)
}
