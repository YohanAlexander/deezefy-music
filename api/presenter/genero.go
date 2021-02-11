package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Genero entidade Genero
type Genero struct {
	Nome     string    `json:"nome"`
	Estilo   string    `json:"estilo"`
	Artistas []Artista `json:"artistas"`
	Musicas  []Musica  `json:"musicas"`
	Perfis   []Perfil  `json:"perfis"`
}

// AppendGenero adiciona presenter na lista
func AppendGenero(genero entity.Genero, generos []Genero) []Genero {
	g := &Genero{}
	g.MakeGenero(genero)
	generos = append(generos, *g)
	return generos
}

// MakeGenero seta os valores a partir da entidade
func (g *Genero) MakeGenero(genero entity.Genero) {
	g.Nome = genero.Nome
	g.Estilo = genero.Estilo
	for _, artista := range genero.Artistas {
		g.Artistas = AppendArtista(artista, g.Artistas)
	}
	for _, musica := range genero.Musicas {
		g.Musicas = AppendMusica(musica, g.Musicas)
	}
	for _, perfil := range genero.Perfis {
		g.Perfis = AppendPerfil(perfil, g.Perfis)
	}
}
