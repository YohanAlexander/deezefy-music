package criaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/criaplaylist"
)

// Service  interface
type Service struct {
	repo Repository
}

// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateCriaPlaylist Create CriaPlaylist
func (s *Service) CreateCriaPlaylist(datacriacao, playlist, usuario string) (string, string, error) {
	e, err := der.NewCriaPlaylist(datacriacao, playlist, usuario)
	if err != nil {
		return err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetCriaPlaylist Get CriaPlaylist
func (s *Service) GetCriaPlaylist(playlist, usuario string) (*der.CriaPlaylist, error) {
	return s.repo.Get(playlist, usuario)
}

// GetCriaPlaylistByPlaylist Get CriaPlaylist By Playlist
func (s *Service) GetCriaPlaylistByPlaylist(playlist string) (*der.CriaPlaylist, error) {
	return s.repo.GetByPlaylist(playlist)
}

// GetCriaPlaylistByUsuario Get CriaPlaylist By Usuario
func (s *Service) GetCriaPlaylistByUsuario(usuario string) (*der.CriaPlaylist, error) {
	return s.repo.GetByUsuario(usuario)
}

// SearchCriaPlaylists Search CriaPlaylists
func (s *Service) SearchCriaPlaylists(query string) ([]*der.CriaPlaylist, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListCriaPlaylists List CriaPlaylists
func (s *Service) ListCriaPlaylists() ([]*der.CriaPlaylist, error) {
	return s.repo.List()
}

// DeleteCriaPlaylist Delete CriaPlaylist
func (s *Service) DeleteCriaPlaylist(playlist, usuario string) error {
	u, err := s.GetCriaPlaylist(playlist, usuario)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(playlist, usuario)
}

// UpdateCriaPlaylist Update CriaPlaylist
func (s *Service) UpdateCriaPlaylist(e *der.CriaPlaylist) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
