package perfilfavoritargenero

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"
)

// Service  interface
type Service struct {
	generoService genero.UseCase
	perfilService perfil.UseCase
}

// NewService create new use case
func NewService(g genero.UseCase, p perfil.UseCase) *Service {
	return &Service{
		generoService: g,
		perfilService: p,
	}
}

// Favoritar favorita um perfil
func (s *Service) Favoritar(g *entity.Genero, p *entity.Perfil) error {
	g, err := s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}
	p, err = s.perfilService.GetPerfil(p.ID)
	if err != nil {
		return err
	}

	err = p.AddGenero(*g)
	if err != nil {
		return err
	}
	err = g.AddPerfil(*p)
	if err != nil {
		return err
	}

	err = s.generoService.UpdateGenero(g)
	if err != nil {
		return err
	}
	err = s.perfilService.UpdatePerfil(p)
	if err != nil {
		return err
	}
	return nil
}

// Desfavoritar desfavorita um perfil
func (s *Service) Desfavoritar(g *entity.Genero, p *entity.Perfil) error {
	g, err := s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}
	p, err = s.perfilService.GetPerfil(p.ID)
	if err != nil {
		return err
	}

	err = p.RemoveGenero(*g)
	if err != nil {
		return err
	}
	err = g.RemovePerfil(*p)
	if err != nil {
		return err
	}

	err = s.generoService.UpdateGenero(g)
	if err != nil {
		return err
	}
	err = s.perfilService.UpdatePerfil(p)
	if err != nil {
		return err
	}
	return nil
}
