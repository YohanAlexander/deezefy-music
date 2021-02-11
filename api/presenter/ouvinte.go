package presenter

import (
	"github.com/yohanalexander/deezefy-music/entity"
)

// Ouvinte entidade Ouvinte
type Ouvinte struct {
	Usuario      Usuario    `json:"usuario"`
	PrimeiroNome string     `json:"primeiro_nome"`
	Sobrenome    string     `json:"sobrenome"`
	Telefones    []string   `json:"telefones"`
	Cria         []Playlist `json:"criadas"`
	Seguindo     []Artista  `json:"seguindo"`
	Curtidas     []Musica   `json:"curtidas"`
	Playlists    []Playlist `json:"playlists"`
	Albums       []Album    `json:"albums"`
}

// AppendOuvinte adiciona presenter na lista
func AppendOuvinte(ouvinte entity.Ouvinte, ouvintes []Ouvinte) []Ouvinte {
	o := &Ouvinte{}
	o.MakeOuvinte(ouvinte)
	ouvintes = append(ouvintes, *o)
	return ouvintes
}

// MakeOuvinte seta os valores a partir da entidade
func (o *Ouvinte) MakeOuvinte(ouvinte entity.Ouvinte) {
	o.Usuario.MakeUsuario(ouvinte.Usuario)
	o.PrimeiroNome = ouvinte.PrimeiroNome
	o.Sobrenome = ouvinte.Sobrenome
	o.Telefones = ouvinte.Telefones
	for _, playlist := range ouvinte.Cria {
		o.Cria = AppendPlaylist(playlist, o.Cria)
	}
	for _, artista := range ouvinte.Seguindo {
		o.Seguindo = AppendArtista(artista, o.Seguindo)
	}
	for _, musica := range ouvinte.Curtidas {
		o.Curtidas = AppendMusica(musica, o.Curtidas)
	}
	for _, playlist := range ouvinte.Playlists {
		o.Playlists = AppendPlaylist(playlist, o.Playlists)
	}
	for _, album := range ouvinte.Albums {
		o.Albums = AppendAlbum(album, o.Albums)
	}
}
