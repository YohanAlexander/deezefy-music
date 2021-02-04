package ouvintesalvarplaylist

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
	"github.com/yohanalexander/deezefy-music/usecase/entity/playlist"
)

// Service  interface
type Service struct {
	ouvinteService  ouvinte.UseCase
	playlistService playlist.UseCase
}

// NewService create new use case
func NewService(o ouvinte.UseCase, p playlist.UseCase) *Service {
	return &Service{
		ouvinteService:  o,
		playlistService: p,
	}
}

// Salvar salva uma playlist
func (s *Service) Salvar(o *entity.Ouvinte, p *entity.Playlist) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	p, err = s.playlistService.GetPlaylist(p.Nome)
	if err != nil {
		return err
	}

	err = p.AddOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.AddPlaylist(*p)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.playlistService.UpdatePlaylist(p)
	if err != nil {
		return err
	}
	return nil
}

// Dessalvar dessalva uma playlist
func (s *Service) Dessalvar(o *entity.Ouvinte, p *entity.Playlist) error {
	o, err := s.ouvinteService.GetOuvinte(o.Usuario.Email)
	if err != nil {
		return err
	}
	p, err = s.playlistService.GetPlaylist(p.Nome)
	if err != nil {
		return err
	}

	err = p.RemoveOuvinte(*o)
	if err != nil {
		return err
	}
	err = o.RemovePlaylist(*p)
	if err != nil {
		return err
	}

	err = s.ouvinteService.UpdateOuvinte(o)
	if err != nil {
		return err
	}
	err = s.playlistService.UpdatePlaylist(p)
	if err != nil {
		return err
	}
	return nil
}
