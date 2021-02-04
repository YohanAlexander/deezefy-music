package albumcontermusica

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Conter(a *entity.Album, m *entity.Musica) error
	Desconter(a *entity.Album, m *entity.Musica) error
}
