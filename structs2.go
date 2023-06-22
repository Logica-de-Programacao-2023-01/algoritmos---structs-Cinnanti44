package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Println("Endereço completo de", p.Nome+":")
	fmt.Println("Rua:", p.Endereco.Rua)
	fmt.Println("Número:", p.Endereco.Numero)
	fmt.Println("Cidade:", p.Endereco.Cidade)
	fmt.Println("Estado:", p.Endereco.Estado)
	fmt.Println()
}

func main() {
	endereco := Endereco{
		Rua:    "Rua A",
		Numero: 123,
		Cidade: "São Paulo",
		Estado: "SP",
	}

	pessoa := Pessoa{
		Nome:     "João",
		Idade:    30,
		Endereco: endereco,
	}

	imprimirEnderecoCompleto(pessoa)
}
