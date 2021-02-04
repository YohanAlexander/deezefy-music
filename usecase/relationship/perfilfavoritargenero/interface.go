package perfilfavoritargenero

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Favoritar(g *entity.Genero, p *entity.Perfil) error
	Desfavoritar(g *entity.Genero, p *entity.Perfil) error
}
