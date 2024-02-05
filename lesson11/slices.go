package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// arr1 := arr[2:5]
	// arr1 := arr[:5]
	// arr1 := arr[3:]
	// arr1 := arr[:]
	// fmt.Println(arr1)

	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// arr1 := arr[:4]
	// fmt.Println(arr1)
	// arr1[0] = 200
	// fmt.Println(arr)
	// fmt.Println(arr1)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr1 := arr[4:]
	fmt.Println(len(arr1))
	fmt.Println(cap(arr))
	fmt.Println(cap(arr1))

	// arr1[0] = 200
	fmt.Println(arr)
	fmt.Println(arr1)

	arr1 = append(arr1, 1000)
	fmt.Println(arr)
	fmt.Println(arr1)

	array := []int{1, 2, 3, 2, 4, 5, 6, 2, 4, 6, 7, 8, 2, 3, 4, 5, 6}
	array2 := make([]int, 10)
	copy(array2, array)
	fmt.Println(array2)

}
