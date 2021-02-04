package ouvintecurtirmusica

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
)

// Service  interface
type Service struct {
	ouvinteService ouvinte.UseCase
	musicaService  musica.UseCase
}

// NewService create new use case
func NewService(o ouvinte.UseCase, m musica.UseCase) *Service {
	return &Service{
		ouvinteService: o,
		musicaService:  m,
	}
}

// Curtir curte uma musica
func (s *Service) Curtir(o *entity.Ouvinte, m *entity.Musica) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.AddOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.AddMusica(*m)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}

// Descurtir descurte uma musica
func (s *Service) Descurtir(o *entity.Ouvinte, m *entity.Musica) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.RemoveOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.RemoveMusica(*m)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}
