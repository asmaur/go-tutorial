package main

import (
	"fmt"
	"math"
)

/*
	interface: regrupamento de comportamento.
				metodos de varios struct tem em comum
	-> para implemetar uma interface,
		uma struct deve todos os metodos da interface
	-> type pode implementar varias interfaces
*/

type forma interface {
	area() float64
}

type figura interface {
	forma
	nome() string
}

type circulo struct {
	raio float64
}

type quadrado struct {
	lado float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

func (c circulo) nome() string {
	return "Circulo"
}

func (q quadrado) nome() string {
	return "Quadrado"
}

func main() {
	// c := circulo{raio: 4}
	//  q := quadrado{lado: 4}

	// formas := []forma{c, q}

	// for _, f := range formas {
	// 	fmt.Printf("Area: %f\n", f.area())
	// }

	//var c figura =

	var q figura = quadrado{lado: 9}
	fmt.Println(q.nome())

}
