package ouvintesalvaralbum

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/album"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
)

// Service  interface
type Service struct {
	ouvinteService ouvinte.UseCase
	albumService   album.UseCase
}

// NewService create new use case
func NewService(o ouvinte.UseCase, a album.UseCase) *Service {
	return &Service{
		ouvinteService: o,
		albumService:   a,
	}
}

// Salvar salva uma album
func (s *Service) Salvar(o *entity.Ouvinte, a *entity.Album) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	a, err = s.albumService.GetAlbum(a.ID)
	if err != nil {
		return err
	}

	err = a.AddOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.AddAlbum(*a)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.albumService.UpdateAlbum(a)
	if err != nil {
		return err
	}
	return nil
}

// Dessalvar dessalva uma album
func (s *Service) Dessalvar(o *entity.Ouvinte, a *entity.Album) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	a, err = s.albumService.GetAlbum(a.ID)
	if err != nil {
		return err
	}

	err = a.RemoveOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.RemoveAlbum(*a)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.albumService.UpdateAlbum(a)
	if err != nil {
		return err
	}
	return nil
}
