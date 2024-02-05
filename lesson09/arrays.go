package main

import "fmt"

/*
	Arrays: usado para guardar dados em memory em sequencia em um unico bloco
	-> cada item é um elemento
	-> os elementos devem ser do mesmo tipo
	-> usar o indice para acessar o dado: iniciando no indice 0
	ex:
		Arrays: |12|1|3|5|9|
		indice: |0 |1|2|3|4|
		tamanho: 5

	Declaração:
	-> var a [5]int // todos os valores serão 0
	-> var a = [5]int{1,2,3,4,5}
	-> var a = [5]int{5, 2:10, 50} // segundo elemento é 10 e o resto 0
	-> var a = [...]int{1, 2, 3, 4}

	-> uma vez definido o tamanho não pode ser alterado. Tamanho fixo. não pode adicionar novo elemento.
*/

func main() {
	// declaração
	// var arr [5]int
	// fmt.Println(arr)

	// var arr = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)

	// arr := [5]int{5, 2: 40, 70}
	// fmt.Println(arr)

	// var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 33, 2, 3}
	// fmt.Println(len(arr))
	// arr[0] = 120
	// fmt.Println(arr)
	// fmt.Println(arr[5])
	// fmt.Println(arr[7])

	arr := [2][2]int{{1, 2}, {5, 6}}
	fmt.Println(arr)
	fmt.Println(arr[1][0])

}
