package main

import "fmt"

/*
	-> acessar os byte de uma string
	-> tentar atualizar um byte de uma string
	-> converter byte e rune para string
	-> fatiar uma string usando slice
*/

func main() {
	// nome := "Telma & Luiza"
	// primeiro := nome[0]
	// ultimo := nome[len(nome)-1]
	// fmt.Println(string(primeiro), string(ultimo))

	// nome[0] = "hjh"

	// var letra rune = 'K'
	// fmt.Println(string(letra))

	// nome := "Telma & Luiza"
	// runeNome := []rune(nome)
	// fmt.Println(runeNome)

	// byteNome := []byte(nome)
	// fmt.Println(byteNome)

	nome := "汉字/漢字"
	fmt.Println(nome)
	primeiro := nome[0]
	fmt.Println(string(primeiro))

	carateres := []rune(nome)
	fmt.Println(carateres)
	fmt.Println(string(carateres[0]))

}
