package main

import "fmt"

/*
	struct: são struturas de dados compostos por dados primitivos.
			criados pelo usuario com diversos campos
	-> definido com: type + nome + struct
	Ex:
		type aluno struct{
			nome string
			sobrenome string
			idade int
			cursos []string / map etc..
		}
	Inicialização:
	-> implicito: sem o nome dos campos
	-> explicito: com nome dos campos
	var aluno1 aluno
	aluno1.nomeDoCampo
*/

type aluno struct {
	nome      string
	sobrenome string
	idade     int
	cursos    []string
}

func main() {
	var aluno1 aluno
	fmt.Printf("%+v\n", aluno1)

	// implicita
	aluno1 = aluno{"andre", "jurandir", 30, []string{"Biologia", "ingles"}}
	fmt.Printf("%+v\n", aluno1)

	// explicita
	aluno2 := aluno{nome: "Jose", sobrenome: "Martins", idade: 45, cursos: []string{"Fisiologia", "Arquitetura"}}
	fmt.Printf("%+v\n", aluno2)

	// struct anonima
	professor := struct {
		nome       string
		disciplina string
	}{
		nome:       "Márcio",
		disciplina: "Arte",
	}

	fmt.Printf("%+v\n", professor)

	fmt.Println(aluno1.nome)
	fmt.Println(professor.disciplina)

}
