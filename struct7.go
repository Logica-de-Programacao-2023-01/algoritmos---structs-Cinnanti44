package main

import "fmt"

type Animal struct {
	nome    string
	espécie string
	idade   int
	som     string
}

func (a *Animal) modificarSom(novoSom string) {
	a.som = novoSom
}

func (a *Animal) imprimirInformações() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Espécie: %s\n", a.espécie)
	fmt.Printf("Idade: %d anos\n", a.idade)
	fmt.Printf("Som: %s\n", a.som)
}

func main() {
	animal := Animal{
		nome:    "Rex",
		espécie: "Cachorro",
		idade:   5,
		som:     "Latido",
	}

	animal.imprimirInformações()

	fmt.Println("Modificando o som...")
	animal.modificarSom("Au Au")

	animal.imprimirInformações()
}
