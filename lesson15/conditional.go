package main

import "fmt"

/*
	condições: if-else, else if -> true/false
*/

func main() {
	idade := 14

	// if idade >= 18 {
	// 	fmt.Println("Você é um adulto")
	// } else if idade > 13 {
	// 	fmt.Println("Você é um adolecente")
	// } else {
	// 	fmt.Println("Você é uma criança!")
	// }

	if num := par(idade); num {
		fmt.Println("Sua idade é par")
	}

}

func par(n int) bool {
	return n&1 == 0
}
