package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// PostMusica entidade Musica
type PostMusica struct {
	ID        int               `json:"id"`
	Nome      string            `json:"nome"`
	Duracao   int               `json:"duracao"`
	Curtiu    []entity.Ouvinte  `json:"curtiu"`
	Artistas  []entity.Artista  `json:"artistas"`
	Playlists []entity.Playlist `json:"playlists"`
	Albums    []entity.Album    `json:"albums"`
	Generos   []entity.Genero   `json:"generos"`
}

// MakeMusica seta os valores a partir da entidade
func (m *PostMusica) MakeMusica(musica entity.Musica) {
	m.ID = musica.ID
	m.Nome = musica.Nome
	m.Duracao = musica.Duracao
	m.Curtiu = musica.Curtiu
	m.Artistas = musica.Gravou
	m.Playlists = musica.Playlists
	m.Albums = musica.Albums
	m.Generos = musica.Generos
}

// Musica presenter Musica
type Musica struct {
	ID        int        `json:"id"`
	Nome      string     `json:"nome"`
	Duracao   int        `json:"duracao"`
	Curtiu    []Ouvinte  `json:"curtiu"`
	Artistas  []Artista  `json:"artistas"`
	Playlists []Playlist `json:"playlists"`
	Albums    []Album    `json:"albums"`
	Generos   []Genero   `json:"generos"`
}

// AppendMusica adiciona presenter na lista
func AppendMusica(musica entity.Musica, musicas []Musica) []Musica {
	m := &Musica{}
	m.MakeMusica(musica)
	musicas = append(musicas, *m)
	return musicas
}

// MakeMusica seta os valores a partir da entidade
func (m *Musica) MakeMusica(musica entity.Musica) {
	m.ID = musica.ID
	m.Nome = musica.Nome
	m.Duracao = musica.Duracao
	for _, ouvinte := range musica.Curtiu {
		m.Curtiu = AppendOuvinte(ouvinte, m.Curtiu)
	}
	for _, artista := range musica.Gravou {
		m.Artistas = AppendArtista(artista, m.Artistas)
	}
	for _, playlist := range musica.Playlists {
		m.Playlists = AppendPlaylist(playlist, m.Playlists)
	}
	for _, album := range musica.Albums {
		m.Albums = AppendAlbum(album, m.Albums)
	}
	for _, genero := range musica.Generos {
		m.Generos = AppendGenero(genero, m.Generos)
	}
}
