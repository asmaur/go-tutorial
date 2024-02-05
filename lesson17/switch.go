package main

import "fmt"

/*
	switch: caso
	valida uma condição dentro outras
*/

func main() {
	val := 8

	switch val {
	case 3:
		fmt.Println("Deu 3")
	case 5:
		fmt.Println("Deu 5")
	case 1, 8, 9:
		fmt.Println("Caso Excepcional!")
	default:
		fmt.Println("caso não encontrado")
	}

}
