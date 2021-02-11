package presenter

import (
	"github.com/yohanalexander/deezefy-music/entity"
)

// Artista entidade Artista
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
