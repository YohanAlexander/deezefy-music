package artistagravarmusica

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
)

// Service  interface
type Service struct {
	artistaService artista.UseCase
	musicaService  musica.UseCase
}

// NewService create new use case
func NewService(a artista.UseCase, m musica.UseCase) *Service {
	return &Service{
		artistaService: a,
		musicaService:  m,
	}
}

// Gravar grava uma musica
func (s *Service) Gravar(a *entity.Artista, m *entity.Musica) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.AddArtista(*a)
	if err != nil {
		return err
	}
	err = a.AddMusica(*m)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}

// Desgravar desgrava uma musica
func (s *Service) Desgravar(a *entity.Artista, m *entity.Musica) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.RemoveArtista(*a)
	if err != nil {
		return err
	}
	err = a.RemoveMusica(*m)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}
