package salvaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaplaylist"
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

// CreateSalvaPlaylist Create SalvaPlaylist
func (s *Service) CreateSalvaPlaylist(playlist, ouvinte string) (string, string, error) {
	e, err := der.NewSalvaPlaylist(playlist, ouvinte)
	if err != nil {
		return err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetSalvaPlaylist Get SalvaPlaylist
func (s *Service) GetSalvaPlaylist(playlist, ouvinte string) (*der.SalvaPlaylist, error) {
	return s.repo.Get(playlist, ouvinte)
}

// GetSalvaPlaylistByPlaylist Get SalvaPlaylist By Playlist
func (s *Service) GetSalvaPlaylistByPlaylist(playlist string) (*der.SalvaPlaylist, error) {
	return s.repo.GetByPlaylist(playlist)
}

// GetSalvaPlaylistByOuvinte Get SalvaPlaylist By Ouvinte
func (s *Service) GetSalvaPlaylistByOuvinte(ouvinte string) (*der.SalvaPlaylist, error) {
	return s.repo.GetByOuvinte(ouvinte)
}

// SearchSalvaPlaylists Search SalvaPlaylists
func (s *Service) SearchSalvaPlaylists(query string) ([]*der.SalvaPlaylist, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListSalvaPlaylists List SalvaPlaylists
func (s *Service) ListSalvaPlaylists() ([]*der.SalvaPlaylist, error) {
	return s.repo.List()
}

// DeleteSalvaPlaylist Delete SalvaPlaylist
func (s *Service) DeleteSalvaPlaylist(playlist, ouvinte string) error {
	u, err := s.GetSalvaPlaylist(playlist, ouvinte)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(playlist, ouvinte)
}

// UpdateSalvaPlaylist Update SalvaPlaylist
func (s *Service) UpdateSalvaPlaylist(e *der.SalvaPlaylist) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
