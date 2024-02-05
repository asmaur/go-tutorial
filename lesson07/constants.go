package main

import "fmt"

/*
	constant : são imutaveis
	palavra chave: const
	-> qualquer dado que não será alterada pode ser declarado como constante
	-> tipos: typed e untyped
	-> declarado a nivel de pacote ou de função
	-> declarado em grupo (mais comum em go)

*/

const (
	x                  = 10
	y          float64 = 60
	appName            = "Curso de Go"
	isRunning          = false
	carater            = 'W'
	verdadeira         = 1 < 9
)

func main() {
	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = y
	fmt.Println(b)

	// const z = a + int(b)
}
