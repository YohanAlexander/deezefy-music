package presenter

import (
	"github.com/yohanalexander/deezefy-music/entity"
)

// Album entidade Album
type Album struct {
	Artista       Artista   `json:"artista"`
	ID            int       `json:"id"`
	Titulo        string    `json:"titulo"`
	AnoLancamento int       `json:"ano_lancamento"`
	Ouvintes      []Ouvinte `json:"ouvintes"`
	Musicas       []Musica  `json:"musicas"`
}

// AppendAlbum adiciona presenter na lista
func AppendAlbum(album entity.Album, albums []Album) []Album {
	a := &Album{}
	a.MakeAlbum(album)
	albums = append(albums, *a)
	return albums
}

// MakeAlbum seta os valores a partir da entidade
func (a *Album) MakeAlbum(album entity.Album) {
	a.Artista.MakeArtista(album.Artista)
	a.ID = album.ID
	a.Titulo = album.Titulo
	a.AnoLancamento = album.AnoLancamento
	for _, ouvinte := range album.Salvou {
		a.Ouvintes = AppendOuvinte(ouvinte, a.Ouvintes)
	}
	for _, musica := range album.Musicas {
		a.Musicas = AppendMusica(musica, a.Musicas)
	}
}
