package artistagravarmusica

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Gravar(a *entity.Artista, m *entity.Musica) error
	Desgravar(a *entity.Artista, m *entity.Musica) error
}
