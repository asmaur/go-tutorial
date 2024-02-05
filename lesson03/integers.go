package main

import "fmt"

/*
	int -> numeros inteiros ex: 1, 122, -23345
	tipos de int: int8, int16, int32, int64 etc -> diferentes tamanhos em memoria.
	unsigned int: uint -> guarda apenas numeros inteiros positivos (uint8, uint16, uint32, uint64)
	byte -> alias para uint8
	rune -> alias para int32
	default: int & uint -> depende da maquina.
*/

func main() {
	// variavel do tipo uint
	var valorPositivo uint8
	valorPositivo = 20
	fmt.Println(valorPositivo)

	// int8
	var valorPositivoNegativo int8
	valorPositivoNegativo = 10
	fmt.Println(valorPositivoNegativo)

	// uint16, unit32 e uint64 -> parecido com uint8
	// int16, int32 e int64 -> parecido com int8

	var valor int = 23456
	fmt.Println(valor)

	var umByte byte = 'B'
	fmt.Println(umByte)

	var umRune rune = 'G'
	fmt.Println(umRune)

	umRune = '🎉'
	fmt.Println(umRune)

}
