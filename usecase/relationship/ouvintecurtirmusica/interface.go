package ouvintecurtirmusica

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Curtir(o *entity.Ouvinte, m *entity.Musica) error
	Descurtir(o *entity.Ouvinte, m *entity.Musica) error
}
