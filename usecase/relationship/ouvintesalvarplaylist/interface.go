package ouvintesalvarplaylist

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Salvar(o *entity.Ouvinte, p *entity.Playlist) error
	Dessalvar(o *entity.Ouvinte, p *entity.Playlist) error
}
