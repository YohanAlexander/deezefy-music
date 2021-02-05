package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Genero entidade Genero
type Genero struct {
	Nome     string           `json:"nome"`
	Estilo   string           `json:"estilo"`
	Artistas []entity.Artista `json:"artistas"`
	Musicas  []entity.Musica  `json:"musicas"`
	Perfis   []entity.Perfil  `json:"perfis"`
}

// GetGenero seta os valores a partir da entidade
func (g *Genero) GetGenero(genero entity.Genero) {
	g.Nome = genero.Nome
	g.Estilo = genero.Estilo
	g.Artistas = genero.Artistas
	g.Musicas = genero.Musicas
	g.Perfis = genero.Perfis
}
