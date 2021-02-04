package ouvintesalvaralbum

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Salvar(o *entity.Ouvinte, a *entity.Album) error
	Dessalvar(o *entity.Ouvinte, a *entity.Album) error
}
