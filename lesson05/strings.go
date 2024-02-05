package main

import "fmt"

/*
	string -> texto: linha simples ou multiplas
	string -> "" ou ``
	string não pode ser alterado diretamente
	rune -> representa um caracter numa string.
*/

func main() {
	var umaString string
	// zero
	fmt.Println(umaString)
	umaString = "Bem vindo na Wanubit!"
	fmt.Println(umaString)

	umaString = "Bem vindo \n na Wanubit"
	fmt.Println(umaString)

	umaString = `Bem vindo ao 
	
	
		Curso de Golang!`
	fmt.Println(umaString)

	nome := "Lazaro"
	sobrenome := "Ramos"
	fmt.Println(nome + " " + sobrenome)

}
