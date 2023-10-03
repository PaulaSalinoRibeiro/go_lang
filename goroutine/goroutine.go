package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Olá mundo!")
	write("Programando em Go lang")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
