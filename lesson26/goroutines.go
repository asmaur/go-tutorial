package main

import (
	"fmt"
	"sync"
)

/*
	goroutine: são funções que rodam em concorrencia/paralello do programa principal
				São thread gerenciado pelo runtime do Go
*/

func responde() {
	fmt.Println("Okay!")
}

var wg = sync.WaitGroup{}

func main() {
	// go responde()

	mensagem := "Valeu!"

	wg.Add(1)
	go func(mensagem string) {
		fmt.Println(mensagem)
		wg.Done()
	}(mensagem)

	mensagem = "De nada!"

	wg.Wait()

	// time.Sleep(100 * time.Millisecond)
}
