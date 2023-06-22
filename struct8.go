package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	origem  string
	destino string
	data    time.Time
	preço   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		// Caso não haja viagens no slice, retorna uma viagem vazia ou trata o erro de outra forma
		return Viagem{}
	}

	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.preço > viagemMaisCara.preço {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: time.Now(), preço: 250.0},
		{origem: "New York", destino: "Los Angeles", data: time.Now(), preço: 350.0},
		{origem: "Paris", destino: "Londres", data: time.Now(), preço: 450.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)
	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.origem)
	fmt.Println("Destino:", viagemMaisCara.destino)
	fmt.Println("Data:", viagemMaisCara.data)
	fmt.Println("Preço:", viagemMaisCara.preço)
}
