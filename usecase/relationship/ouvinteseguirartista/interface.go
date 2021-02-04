package ouvinteseguirartista

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Seguir(o *entity.Ouvinte, a *entity.Artista) error
	Desseguir(o *entity.Ouvinte, a *entity.Artista) error
}
