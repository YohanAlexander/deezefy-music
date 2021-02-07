package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Playlist entidade Playlist
type Playlist struct {
	Nome           string           `json:"nome"`
	Status         string           `json:"status"`
	DataCriacao    string           `json:"datacriacao"`
	Ouvintes       []entity.Ouvinte `json:"ouvintes"`
	Musicas        []entity.Musica  `json:"musicas"`
	NumeroOuvintes int              `json:"numeroouvintes"`
}

// GetPlaylist seta os valores a partir da entidade
func (p *Playlist) GetPlaylist(playlist entity.Playlist) {
	p.Nome = playlist.Nome
	p.Status = playlist.Status
	p.DataCriacao = playlist.DataCriacao
	p.Ouvintes = playlist.Salvou
	p.Musicas = playlist.Musicas
	p.NumeroOuvintes = len(p.Ouvintes)
}
