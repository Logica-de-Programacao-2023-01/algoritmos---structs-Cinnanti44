package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(index int) {
	if index >= 0 && index < len(a.notas) {
		a.notas = append(a.notas[:index], a.notas[index+1:]...)
	}
}

func (a *Aluno) calcularMédia() float64 {
	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	if len(a.notas) > 0 {
		return total / float64(len(a.notas))
	}
	return 0.0
}

func (a *Aluno) imprimirInformações() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Média: %.2f\n", a.calcularMédia())
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{8.5, 9.0, 7.5},
	}

	aluno.imprimirInformações()

	aluno.adicionarNota(8.0)
	aluno.removerNota(1)

	aluno.imprimirInformações()
}
