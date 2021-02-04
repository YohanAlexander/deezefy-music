package playlistcontermusica

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
	"github.com/yohanalexander/deezefy-music/usecase/entity/playlist"
)

// Service  interface
type Service struct {
	playlistService playlist.UseCase
	musicaService   musica.UseCase
}

// NewService create new use case
func NewService(p playlist.UseCase, m musica.UseCase) *Service {
	return &Service{
		playlistService: p,
		musicaService:   m,
	}
}

// Conter conter uma musica
func (s *Service) Conter(p *entity.Playlist, m *entity.Musica) error {
	p, err := s.playlistService.GetPlaylist(p.Nome)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.AddPlaylist(*p)
	if err != nil {
		return err
	}
	err = p.AddMusica(*m)
	if err != nil {
		return err
	}

	err = s.playlistService.UpdatePlaylist(p)
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
func (s *Service) Desconter(p *entity.Playlist, m *entity.Musica) error {
	p, err := s.playlistService.GetPlaylist(p.Nome)
	if err != nil {
		return err
	}
	m, err = s.musicaService.GetMusica(m.ID)
	if err != nil {
		return err
	}

	err = m.RemovePlaylist(*p)
	if err != nil {
		return err
	}
	err = p.RemoveMusica(*m)
	if err != nil {
		return err
	}

	err = s.playlistService.UpdatePlaylist(p)
	if err != nil {
		return err
	}
	err = s.musicaService.UpdateMusica(m)
	if err != nil {
		return err
	}
	return nil
}
