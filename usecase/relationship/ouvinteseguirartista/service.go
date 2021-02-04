package ouvinteseguirartista

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
)

// Service  interface
type Service struct {
	ouvinteService ouvinte.UseCase
	artistaService artista.UseCase
}

// NewService create new use case
func NewService(o ouvinte.UseCase, a artista.UseCase) *Service {
	return &Service{
		ouvinteService: o,
		artistaService: a,
	}
}

// Seguir segue um artista
func (s *Service) Seguir(o *entity.Ouvinte, a *entity.Artista) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	a, err = s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}

	err = a.AddOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.AddArtista(*a)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	return nil
}

// Desseguir dessegue um artista
func (s *Service) Desseguir(o *entity.Ouvinte, a *entity.Artista) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	a, err = s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}

	err = a.RemoveOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.RemoveArtista(*a)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	return nil
}
