package playlistcontermusica

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Conter(p *entity.Playlist, m *entity.Musica) error
	Desconter(p *entity.Playlist, m *entity.Musica) error
}
