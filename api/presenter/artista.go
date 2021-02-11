package presenter

import (
	"github.com/yohanalexander/deezefy-music/entity"
)

// PostArtista entidade Artista
type PostArtista struct {
	Usuario       entity.Usuario   `json:"usuario"`
	NomeArtistico string           `json:"nome_artistico"`
	Biografia     string           `json:"biografia"`
	AnoFormacao   int              `json:"ano_formacao"`
	Organizador   []entity.Evento  `json:"eventos"`
	Seguidores    []entity.Ouvinte `json:"seguidores"`
	Musicas       []entity.Musica  `json:"musicas"`
	Perfis        []entity.Perfil  `json:"perfis"`
	Generos       []entity.Genero  `json:"generos"`
	Albums        []entity.Album   `json:"albums"`
}

// MakeArtista seta os valores a partir da entidade
func (a *PostArtista) MakeArtista(artista entity.Artista) {
	a.Usuario = artista.Usuario
	a.NomeArtistico = artista.NomeArtistico
	a.Biografia = artista.Biografia
	a.AnoFormacao = artista.AnoFormacao
	a.Organizador = artista.Organizador
	a.Seguidores = artista.Seguidores
	a.Musicas = artista.Grava
	a.Perfis = artista.Perfis
	a.Generos = artista.Generos
	a.Albums = artista.Albums
}

// Artista presenter Artista
type Artista struct {
	Usuario       Usuario   `json:"usuario"`
	NomeArtistico string    `json:"nome_artistico"`
	Biografia     string    `json:"biografia"`
	AnoFormacao   int       `json:"ano_formacao"`
	Organizador   []Evento  `json:"eventos"`
	Seguidores    []Ouvinte `json:"seguidores"`
	Musicas       []Musica  `json:"musicas"`
	Perfis        []Perfil  `json:"perfis"`
	Generos       []Genero  `json:"generos"`
	Albums        []Album   `json:"albums"`
}

// AppendArtista adiciona presenter na lista
func AppendArtista(artista entity.Artista, artistas []Artista) []Artista {
	a := &Artista{}
	a.MakeArtista(artista)
	artistas = append(artistas, *a)
	return artistas
}

// MakeArtista seta os valores a partir da entidade
func (a *Artista) MakeArtista(artista entity.Artista) {
	a.Usuario.MakeUsuario(artista.Usuario)
	a.NomeArtistico = artista.NomeArtistico
	a.Biografia = artista.Biografia
	a.AnoFormacao = artista.AnoFormacao
	for _, evento := range artista.Organizador {
		a.Organizador = AppendEvento(evento, a.Organizador)
	}
	for _, ouvinte := range artista.Seguidores {
		a.Seguidores = AppendOuvinte(ouvinte, a.Seguidores)
	}
	for _, musica := range artista.Grava {
		a.Musicas = AppendMusica(musica, a.Musicas)
	}
	for _, perfil := range artista.Perfis {
		a.Perfis = AppendPerfil(perfil, a.Perfis)
	}
	for _, genero := range artista.Generos {
		a.Generos = AppendGenero(genero, a.Generos)
	}
	for _, album := range artista.Albums {
		a.Albums = AppendAlbum(album, a.Albums)
	}
}
