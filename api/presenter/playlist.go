package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Playlist entidade Playlist
type Playlist struct {
	Usuario        Usuario   `json:"usuario"`
	Nome           string    `json:"nome"`
	Status         string    `json:"status"`
	DataCriacao    string    `json:"data_criacao"`
	NumeroOuvintes int       `json:"numero_ouvintes"`
	Ouvintes       []Ouvinte `json:"ouvintes"`
	Musicas        []Musica  `json:"musicas"`
}

// AppendPlaylist adiciona presenter na lista
func AppendPlaylist(playlist entity.Playlist, playlists []Playlist) []Playlist {
	p := &Playlist{}
	p.MakePlaylist(playlist)
	playlists = append(playlists, *p)
	return playlists
}

// MakePlaylist seta os valores a partir da entidade
func (p *Playlist) MakePlaylist(playlist entity.Playlist) {
	p.Usuario.MakeUsuario(playlist.Usuario)
	p.Nome = playlist.Nome
	p.Status = playlist.Status
	p.DataCriacao = playlist.DataCriacao
	p.NumeroOuvintes = len(p.Ouvintes)
	for _, ouvinte := range playlist.Salvou {
		p.Ouvintes = AppendOuvinte(ouvinte, p.Ouvintes)
	}
	for _, musica := range playlist.Musicas {
		p.Musicas = AppendMusica(musica, p.Musicas)
	}
}
