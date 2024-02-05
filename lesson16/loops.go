package main

import "fmt"

/*
	loop: repete um bloco de codigo enquanto uma condição está verdadeira.
		: no caso um determinado numero de vezes
	-> 3 tipos de for loop:
			-> for c/ condição
			-> for c/ clausulas
			-> for c/ faixa de iteração
	-> break : para finalizar um loop de forma abrupta
*/

func main() {
	// a := 10

	// // while loop
	// for {
	// 	fmt.Println(a)
	// 	a--
	// 	if a == 0 {
	// 		break
	// 	}
	// }

	// for i := 0; i <= 20; i++ {
	// 	fmt.Println(i)
	// }

	// arr := []int{11, 29, 37, 47, 51, 63}

	// for index, valor := range arr {
	// 	fmt.Printf("indice: %d, valor: %d\n", index, valor)
	// }

	maps := map[string]int{"a": 2, "b": 45, "t": 50}
	// for chave, valor := range maps {
	// 	fmt.Printf("chave: %s, valor: %d\n", chave, valor)
	// }

	for _, valor := range maps {
		fmt.Printf("valor: %d\n", valor)
	}

}
