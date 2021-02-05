package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Album entidade Album
type Album struct {
	Artista       entity.Artista   `json:"artista"`
	ID            int              `json:"id"`
	Titulo        string           `json:"titulo"`
	AnoLancamento int              `json:"anolancamento"`
	Ouvintes      []entity.Ouvinte `json:"ouvintes"`
	Musicas       []entity.Musica  `json:"musicas"`
}

// GetAlbum seta os valores a partir da entidade
func (a *Album) GetAlbum(album entity.Album) {
	a.Artista = album.Artista
	a.ID = album.ID
	a.Titulo = album.Titulo
	a.AnoLancamento = album.AnoLancamento
	a.Ouvintes = album.Salvou
	a.Musicas = album.Musicas
}
