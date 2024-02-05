package main

import (
	"fmt"
)

/*
	maps: chave-valor.
	clientes := map[string]int{
		"nome": Aline,
		"sobrenome": "Prates"
	}
	-> zero: nil
		Ex:
			nome := map[string]int{} // vazia
			nome := make(map[string]int) // nil
	-> clientes["nome"] : retorna o valor, se não existir o zero default é retornado
	-> dado, ok := clientes["nome"] // o ok=true/false se a chave existir.
	-> adicionar novo valor: cliente["alguem"] = "flavio"

*/

func main() {
	// map nil
	// var numeros map[string]int
	// fmt.Println(numeros == nil)
	// fmt.Println(len(numeros))
	// fmt.Println(numeros["algo"])
	// numeros["nome"] = 56
	// fmt.Println(numeros["nome"])

	// declaração
	// numeros := map[string]int{}

	// usando a função make
	// numeros := make(map[string]int)

	// declaração literal
	// var numeros map[string]int = map[string]int{}

	numeros := map[string]int{
		"quatro": 5,
	}
	numeros["um"] = 10
	numeros["dois"] = 100
	// fmt.Println(numeros)

	// fmt.Println(numeros["um"])

	// retorno ok
	numeros["x"] = 0
	valor, ok := numeros["xdfhfghfgh"]
	fmt.Println(valor, ok)

}
