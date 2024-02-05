package main

import "fmt"

/*
Os tipos primitivos:
- Boleano: true ou false (verdadeiro ou falso) representa um determinado estado
- String: são conjunto de caracteres que formam textos
- Numeros: int, float e complexos.
- o conceito de valor zero
*/
func main() {
	var minhaString string
	// zero para um string ""
	fmt.Println(minhaString)
	minhaString = "Bem-vindo na Wanubit"
	fmt.Println(minhaString)

	var meuInt int
	// o zero para inteiro é o 0
	fmt.Println(meuInt)
	meuInt = 20
	fmt.Println(meuInt)

	var minhaFloat float32
	fmt.Println(minhaFloat)
	minhaFloat = 3.14
	fmt.Println(minhaFloat)

	var meuBoleano bool
	fmt.Println(meuBoleano)
	meuBoleano = true
	fmt.Println(meuBoleano)

	minhaString = "Obrigado por estarem aqui!"
	fmt.Println(minhaString)
}
