package main

import "fmt"

type Musica struct {
	titulo  string
	artista string
	duracao int // em segundos
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Nome da playlist: %s\n", p.nome)

	duracaoTotal := 0
	for _, musica := range p.musicas {
		fmt.Printf("Título: %s\n", musica.titulo)
		fmt.Printf("Artista: %s\n", musica.artista)
		fmt.Printf("Duração: %d segundos\n", musica.duracao)
		fmt.Println()
		duracaoTotal += musica.duracao
	}

	fmt.Printf("Tempo total da playlist: %d segundos\n", duracaoTotal)
}

func main() {
	// Exemplo de uso
	playlist := Playlist{
		nome: "Minha Playlist",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			{titulo: "Música 2", artista: "Artista 2", duracao: 240},
			{titulo: "Música 3", artista: "Artista 3", duracao: 200},
		},
	}

	imprimirPlaylist(playlist)
}
