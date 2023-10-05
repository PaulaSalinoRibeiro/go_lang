package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // declaração de um canal

	go escrever("Olá mundo", canal)

	for {
		mensagem, aberto := <-canal // canal envia o valor
		fmt.Println(aberto)
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text // canal recebe o valor de text
		time.Sleep(time.Second)
	}
	close(canal)
}
