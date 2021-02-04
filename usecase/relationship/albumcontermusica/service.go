package albumcontermusica

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
	"github.com/yohanalexander/deezefy-music/usecase/entity/album"
)

// Service  interface
type Service struct {
	albumService album.UseCase
	musicaService   musica.UseCase
}

// NewService create new use case
func NewService(a album.UseCase, m musica.UseCase) *Service {
	return &Service{
		albumService: a,
		musicaService:   m,
	}
}

// Conter conter uma musica
func (s *Service) Conter(a *entity.Album, m *entity.Musica) error {
	a, err := s.albumService.GetAlbum(a.ID)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.AddAlbum(*a)
	if err != nil {
		return err
	}
	err = a.AddMusica(*m)
	if err != nil {
		return err
	}

	err = s.albumService.UpdateAlbum(a)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}

// Desconter desconter uma musica
func (s *Service) Desconter(a *entity.Album, m *entity.Musica) error {
	a, err := s.albumService.GetAlbum(a.ID)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.RemoveAlbum(*a)
	if err != nil {
		return err
	}
	err = a.RemoveMusica(*m)
	if err != nil {
		return err
	}

	err = s.albumService.UpdateAlbum(a)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}
