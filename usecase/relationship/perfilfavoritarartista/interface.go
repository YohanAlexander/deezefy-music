package perfilavoritarartista

import "github.com/yohanalexander/deezefy-music/entity"

// UseCase interface
type UseCase interface {
	Favoritar(a *entity.Artista, p *entity.Perfil) error
	Desfavoritar(a *entity.Artista, p *entity.Perfil) error
}
