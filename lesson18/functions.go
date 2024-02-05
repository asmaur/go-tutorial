package main

import "fmt"

/*
	função: bloco de código que pode ser chamado para executar uma tarefa.
	-> Pode ou não retornar um valor
	-> pode ter parametro ou não.
*/

func teste(x string) {
	fmt.Println("Sou uma função: ", x)
}

func soma(x, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

func main() {
	teste("Go")
	// fmt.Println(soma(2, 3))
	z, w := soma(10, 5)
	fmt.Println(z, w)
}
