package main

import "fmt"

/*
	Similiar a classes em outras linguagens
*/

type ponto struct {
	x float32
	y float32
}

type circulo struct {
	raio float32
	*ponto
}

func mudarPonto(pt *ponto, a float32) {
	pt.x = a
}

func main() {
	// p1 := ponto{x: 2, y: 3}
	// fmt.Println(p1)

	p1 := &ponto{x: 7, y: 19}
	p1.y = 200
	// fmt.Println(p1.y)

	mudarPonto(p1, 90)
	// fmt.Println(p1.x)

	c1 := circulo{3.14, p1}
	fmt.Println(c1.y)
}
