package main

import "fmt"

/*
	-> função pode agir como uma variavel
	-> pode retornar uma função
	-> funções anonimas

*/

func simples() {
	fmt.Println("Oi")
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	x := simples
	x()

	y := func() { fmt.Println("Sou uma função em linha!") }
	y()

	z := func(d int) int {
		return d * -2
	}(20)
	fmt.Println(z)

	returnFunc("Golang")()

	r := returnFunc("Retornei")
	r()

}
