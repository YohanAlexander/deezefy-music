package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Musica entidade Musica
type Musica struct {
	ID        int               `json:"id"`
	Nome      string            `json:"nome"`
	Duracao   int               `json:"duracao"`
	Curtiu    []entity.Ouvinte  `json:"curtiu"`
	Artistas  []entity.Artista  `json:"artistas"`
	Playlists []entity.Playlist `json:"playlists"`
	Albums    []entity.Album    `json:"albums"`
	Generos   []entity.Genero   `json:"generos"`
}

// GetMusica seta os valores a partir da entidade
func (m *Musica) GetMusica(musica entity.Musica) {
	m.ID = musica.ID
	m.Nome = musica.Nome
	m.Duracao = musica.Duracao
	m.Curtiu = musica.Curtiu
	m.Artistas = musica.Gravou
	m.Playlists = musica.Playlists
	m.Albums = musica.Albums
	m.Generos = musica.Generos
}
