package salvaalbum

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaalbum"
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

// CreateSalvaAlbum Create SalvaAlbum
func (s *Service) CreateSalvaAlbum(album int, ouvinte, artista string) (int, string, string, error) {
	e, err := der.NewSalvaAlbum(album, ouvinte, artista)
	if err != nil {
		return 1, err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetSalvaAlbum Get SalvaAlbum
func (s *Service) GetSalvaAlbum(album int, ouvinte, artista string) (*der.SalvaAlbum, error) {
	return s.repo.Get(album, ouvinte, artista)
}

// GetSalvaAlbumByAlbum Get SalvaAlbum By Album
func (s *Service) GetSalvaAlbumByAlbum(album int) (*der.SalvaAlbum, error) {
	return s.repo.GetByAlbum(album)
}

// GetSalvaAlbumByOuvinte Get SalvaAlbum By Ouvinte
func (s *Service) GetSalvaAlbumByOuvinte(ouvinte string) (*der.SalvaAlbum, error) {
	return s.repo.GetByOuvinte(ouvinte)
}

// SearchSalvaAlbums Search SalvaAlbums
func (s *Service) SearchSalvaAlbums(query string) ([]*der.SalvaAlbum, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListSalvaAlbums List SalvaAlbums
func (s *Service) ListSalvaAlbums() ([]*der.SalvaAlbum, error) {
	return s.repo.List()
}

// DeleteSalvaAlbum Delete SalvaAlbum
func (s *Service) DeleteSalvaAlbum(album int, ouvinte, artista string) error {
	u, err := s.GetSalvaAlbum(album, ouvinte, artista)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(album, ouvinte, artista)
}

// UpdateSalvaAlbum Update SalvaAlbum
func (s *Service) UpdateSalvaAlbum(e *der.SalvaAlbum) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
