package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // quantidade de goroutine

	go func() {
		write("Olá mundo!")
		waitGroup.Done() // inidca que foi concluida
	}()

	go func() {
		write("Programando em Go lang")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
