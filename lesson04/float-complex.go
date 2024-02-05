package main

import "fmt"

/*
	float -> ponto flutuante: não são preciso (exatos)
	temos: float32 e float64 -> default é float64

	complexos -> real e imag
	temos: complex64 e complex128
*/

func main() {
	var valorFloat float32
	fmt.Println(valorFloat)
	valorFloat = 23.4456567
	fmt.Println(valorFloat)

	var valorFloat64 float64
	fmt.Println(valorFloat64)
	valorFloat64 = 456.45656854646456878787
	fmt.Println(valorFloat64)

	var valorComplexo complex128
	valorComplexo = complex(valorFloat64, valorFloat64)
	fmt.Println(valorComplexo)

	var parteReal, parteImag float64
	parteReal = real(valorComplexo)
	parteImag = imag(valorComplexo)
	fmt.Println(parteReal, parteImag)
}
