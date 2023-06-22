package main

import (
	"fmt"
	"time"
)

type Funcionário struct {
	nome    string
	salário float64
	idade   int
}

func (f *Funcionário) aumentarSalário(porcentagem float64) {
	aumento := f.salário * (porcentagem / 100)
	f.salário += aumento
}

func (f *Funcionário) diminuirSalário(porcentagem float64) {
	diminuição := f.salário * (porcentagem / 100)
	f.salário -= diminuição
}

func (f *Funcionário) tempoServiço() int {
	idadeComeçoTrabalho := 18
	anosServiço := time.Now().Year() - (f.idade + idadeComeçoTrabalho)
	return anosServiço
}

func main() {
	funcionário := Funcionário{
		nome:    "João",
		salário: 2000,
		idade:   25,
	}

	fmt.Println("Salário antes do aumento:", funcionário.salário)
	funcionário.aumentarSalário(10)
	fmt.Println("Salário após o aumento:", funcionário.salário)

	fmt.Println("Salário antes da diminuição:", funcionário.salário)
	funcionário.diminuirSalário(5)
	fmt.Println("Salário após a diminuição:", funcionário.salário)

	fmt.Println("Tempo de serviço:", funcionário.tempoServiço(), "anos")
}
