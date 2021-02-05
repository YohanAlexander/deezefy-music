package musicapossuirgenero

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
)

// Service  interface
type Service struct {
	musicaService musica.UseCase
	generoService genero.UseCase
}

// NewService create new use case
func NewService(a musica.UseCase, g genero.UseCase) *Service {
	return &Service{
		musicaService: a,
		generoService: g,
	}
}

// Possuir possuir um genero
func (s *Service) Possuir(a *entity.Musica, g *entity.Genero) error {
	a, err := s.musicaService.GetMusica(a.ID)
	if err != nil {
		return err
	}
	g, err = s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}

	err = g.AddMusica(*a)
	if err != nil {
		return err
	}
	err = a.AddGenero(*g)
	if err != nil {
		return err
	}

	err = s.musicaService.UpdateMusica(a)
	if err != nil {
		return err
	}
	err = s.generoService.UpdateGenero(g)
	if err != nil {
		return err
	}
	return nil
}

// Despossuir despossuir um genero
func (s *Service) Despossuir(a *entity.Musica, g *entity.Genero) error {
	a, err := s.musicaService.GetMusica(a.ID)
	if err != nil {
		return err
	}
	g, err = s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}

	err = g.RemoveMusica(*a)
	if err != nil {
		return err
	}
	err = a.RemoveGenero(*g)
	if err != nil {
		return err
	}

	err = s.musicaService.UpdateMusica(a)
	if err != nil {
		return err
	}
	err = s.generoService.UpdateGenero(g)
	if err != nil {
		return err
	}
	return nil
}
