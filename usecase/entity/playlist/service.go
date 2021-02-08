package playlist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
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
func (s *Service) CreatePlaylist(email, password, birthday, nome, status, datacriacao string) (string, error) {
	e, err := entity.NewPlaylist(email, password, birthday, nome, status, datacriacao)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetPlaylist Get Playlist
func (s *Service) GetPlaylist(nome string) (*entity.Playlist, error) {
	return s.repo.Get(nome)
}

// SearchPlaylists Search Playlists
func (s *Service) SearchPlaylists(query string) ([]*entity.Playlist, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListPlaylists List Playlists
func (s *Service) ListPlaylists() ([]*entity.Playlist, error) {
	return s.repo.List()
}

// DeletePlaylist Delete Playlist
func (s *Service) DeletePlaylist(nome string) error {
	u, err := s.GetPlaylist(nome)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(nome)
}

// UpdatePlaylist Update Playlist
func (s *Service) UpdatePlaylist(e *entity.Playlist) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
