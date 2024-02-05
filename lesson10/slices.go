package main

import "fmt"

/*
	Slice: são array sem tamanho fixo, pode se adicionar quanto elementos desejar
	-> não tem limitações de um array
	-> len: tamanho do slice ou array
	-> capacidade: quanto elemento pode conter seu slice no momento (cap)
	-> |3|5|6|8| | | -> len = 4 e cap = 6
	Decl:
	-> var arr []int
	-> arr := []int{1,2,3,4}
	-> arr := []int{34, 2:9, 5}
	-> arr := make([]int, 0) | make([]string, 5)
	-> append: adiciona um elemento ao slice
		Ex:
			arr = append(arr, 5)
			arr = append(a, 1,2,3,4)
			arr = append(a, b...) // b sendo um outro slice
	-> var arr int : nil (null)
	-> var arr := []int{} : slice vazia
*/

func main() {
	// slice nil
	var arr1 []int
	fmt.Println(arr1 == nil)

	// slice vazia
	var arr = []int{}
	fmt.Println(arr == nil)

	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr2)

	arr3 := make([]string, 0)
	fmt.Println(arr3)
	arr3 = append(arr3, "Edu")
	fmt.Println(arr3)
	arr3 = append(arr3, "Maria")
	fmt.Println(arr3)

	arr4 := []int{1, 2, 3, 4}
	arr4 = append(arr4, 34, 56, 78)
	fmt.Println(arr4)

	b := []int{90, 89, 67}
	arr4 = append(arr4, b...)
	fmt.Println(arr4)

}
