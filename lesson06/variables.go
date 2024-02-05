package main

import "fmt"

/*
	variaveis: são usados para guardar dados que podem mudar ao longo do programa.
	-> variaveis de nível de pacote ou a nível de função
	-> declaração explicita ou implicita do tipo da variável
	-> declaração em grupo
	-> multiplas variaveis podem ser declados em uma unica linha.
	Ex:
	var idade int
	var idade int = 10
	var idade = 10
	var idade, nome = 25, "Leandro"
	idade := 25
	idade, nome := 30, "Edu"
*/

// var (
// 	umaInt    int    = 35
// 	umaString string = "Wanubit"
// )

func main() {
	var umaInt int = 35
	var umaString string = "Wanubit"
	fmt.Println(umaInt)
	fmt.Println(umaString)

	// implicita
	var idade = 25
	fmt.Println(idade)

	// muliplas variaveis
	var ano, mes = 2024, "Janeiro"
	fmt.Println(ano, mes)

	// declarações curtas
	pais := "Brazil"
	cidade := "São Paulo"
	fmt.Println(pais, cidade)

	var numero int // zero = 0
	var homens int = 300
	var numeroPI = 3.14
	fmt.Println(numero, homens, numeroPI)

	estado, cidade := "Rio de Janeiro", "Rio de Janeiro"
	fmt.Println(estado, cidade)

}
