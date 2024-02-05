package main

import "fmt"

/*
	pointer: serve para acessar referencia de variaveis
	-> & operador de referenciação: criar um pointer (endereço)
	-> * operador de dereferencia: acessa o valor armazenado na memoria
		(valor)
*/

func mudar(str *string) {
	*str = "choquei!"
}

func main() {
	// x := 9
	// y := &x

	// fmt.Println(x, *y)
	// *y = 10
	// fmt.Println(x, y)
	// fmt.Println(x, *y)

	mudei := "popo"
	fmt.Println(mudei)
	mudar(&mudei)
	fmt.Println(mudei)

	var pointer *string = &mudei
	fmt.Println(pointer, &pointer)

}
