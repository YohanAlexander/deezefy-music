package playlist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/playlist"
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

// CreatePlaylist Create Playlist
func (s *Service) CreatePlaylist(nome, status string) (string, error) {
	e, err := der.NewPlaylist(nome, status)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetPlaylist Get Playlist
func (s *Service) GetPlaylist(email string) (*der.Playlist, error) {
	return s.repo.Get(email)
}

// SearchPlaylists Search Playlists
func (s *Service) SearchPlaylists(query string) ([]*der.Playlist, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListPlaylists List Playlists
func (s *Service) ListPlaylists() ([]*der.Playlist, error) {
	return s.repo.List()
}

// DeletePlaylist Delete Playlist
func (s *Service) DeletePlaylist(email string) error {
	u, err := s.GetPlaylist(email)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(email)
}

// UpdatePlaylist Update Playlist
func (s *Service) UpdatePlaylist(e *der.Playlist) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}