package main

import "fmt"

type Filme struct {
	título     string
	diretor    string
	ano        int
	avaliações []int
}

func (f *Filme) adicionarAvaliação(avaliação int) {
	f.avaliações = append(f.avaliações, avaliação)
}

func (f *Filme) removerAvaliação(index int) {
	if index >= 0 && index < len(f.avaliações) {
		f.avaliações = append(f.avaliações[:index], f.avaliações[index+1:]...)
	}
}

func (f *Filme) calcularMédiaAvaliações() float64 {
	total := 0
	for _, avaliação := range f.avaliações {
		total += avaliação
	}
	if len(f.avaliações) > 0 {
		return float64(total) / float64(len(f.avaliações))
	}
	return 0.0
}

func (f *Filme) imprimirInformações() {
	fmt.Printf("Título: %s\n", f.título)
	fmt.Printf("Diretor: %s\n", f.diretor)
	fmt.Printf("Ano: %d\n", f.ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.calcularMédiaAvaliações())
}

func main() {
	filme := Filme{
		título:     "Inception",
		diretor:    "Christopher Nolan",
		ano:        2010,
		avaliações: []int{9, 8, 9, 10},
	}

	filme.imprimirInformações()

	filme.adicionarAvaliação(7)
	filme.removerAvaliação(1)

	filme.imprimirInformações()
}
