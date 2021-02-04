package artistapossuirgenero

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Possuir(a *entity.Artista, g *entity.Genero) error
	Despossuir(a *entity.Artista, g *entity.Genero) error
}
