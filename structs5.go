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

func buscarPorTitulo(titulo string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	// Exemplo de uso
	playlists := []Playlist{
		{
			nome: "Playlist 1",
			musicas: []Musica{
				{titulo: "Música 1", artista: "Artista 1", duracao: 180},
				{titulo: "Música 2", artista: "Artista 2", duracao: 240},
				{titulo: "Música 3", artista: "Artista 3", duracao: 200},
			},
		},
		{
			nome: "Playlist 2",
			musicas: []Musica{
				{titulo: "Música 4", artista: "Artista 4", duracao: 210},
				{titulo: "Música 5", artista: "Artista 5", duracao: 190},
			},
		},
		{
			nome: "Playlist 3",
			musicas: []Musica{
				{titulo: "Música 6", artista: "Artista 6", duracao: 220},
				{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			},
		},
	}

	tituloBuscado := "Música 1"
	resultado := buscarPorTitulo(tituloBuscado, playlists)

	if len(resultado) > 0 {
		fmt.Printf("Playlists que possuem a música '%s':\n", tituloBuscado)
		for _, playlist := range resultado {
			fmt.Println(playlist.nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist possui a música '%s'.\n", tituloBuscado)
	}
}
