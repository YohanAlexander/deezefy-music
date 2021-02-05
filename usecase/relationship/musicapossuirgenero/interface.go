package musicapossuirgenero

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Possuir(m *entity.Musica, g *entity.Genero) error
	Despossuir(m *entity.Musica, g *entity.Genero) error
}
