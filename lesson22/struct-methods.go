package main

import "fmt"

/*
	struct sendo similar a classes
	-> podem ter metodos que especificam suas hablidades.
	-> seja se comportam como classes (com atributos e metodos)
*/

type aluno struct {
	nome  string
	idade int
	notas []int
}

type ponto struct {
	x int
}

func (a aluno) getNome() string {
	return a.nome
}

func (a *aluno) setIdade(idade int) {
	a.idade = idade
}

func (a aluno) getIdade() int {
	return a.idade
}

func (a aluno) getMedia() float32 {
	soma := 0
	for _, v := range a.notas {
		soma += v
	}
	return float32(soma) / float32(len(a.notas))
}

func (a aluno) getMaxNota() int {
	max := 0

	for _, v := range a.notas {
		if max < v {
			max = v
		}
	}
	return max
}

func (p ponto) getX() int {
	return p.x
}

func main() {
	aluno1 := aluno{nome: "Satoro", idade: 18, notas: []int{5, 7, 8}}
	aluno2 := aluno{nome: "Sukuna", idade: 200, notas: []int{7, 8, 9, 10}}
	fmt.Println(aluno1.getNome())
	fmt.Println(aluno2.getNome())

	aluno2.setIdade(20)
	fmt.Println(aluno2.getIdade())

	fmt.Println(aluno1.getMedia())
	fmt.Println(aluno2.getMedia())

	fmt.Printf("Nota máxima do aluno: %s é %d\n", aluno1.getNome(), aluno1.getMaxNota())
	fmt.Printf("Nota máxima do aluno: %s é %d\n", aluno2.getNome(), aluno2.getMaxNota())

	p1 := ponto{x: 8}
	fmt.Println(p1.getX())

}
