package artistapossuirgenero

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"
)

// Service  interface
type Service struct {
	artistaService artista.UseCase
	generoService  genero.UseCase
}

// NewService create new use case
func NewService(a artista.UseCase, g genero.UseCase) *Service {
	return &Service{
		artistaService: a,
		generoService:  g,
	}
}

// Possuir possuir um genero
func (s *Service) Possuir(a *entity.Artista, g *entity.Genero) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	g, err = s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}

	err = g.AddArtista(*a)
	if err != nil {
		return err
	}
	err = a.AddGenero(*g)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
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
func (s *Service) Despossuir(a *entity.Artista, g *entity.Genero) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	g, err = s.generoService.GetGenero(g.Nome)
	if err != nil {
		return err
	}

	err = g.RemoveArtista(*a)
	if err != nil {
		return err
	}
	err = a.RemoveGenero(*g)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	err = s.generoService.UpdateGenero(g)
	if err != nil {
		return err
	}
	return nil
}
