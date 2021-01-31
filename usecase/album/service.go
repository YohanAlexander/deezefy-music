package album

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/album"
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

// CreateAlbum Create Album
func (s *Service) CreateAlbum(id, anolancamento int, titulo, artista string) (int, error) {
	e, err := der.NewAlbum(id, anolancamento, titulo, artista)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetAlbum Get Album
func (s *Service) GetAlbum(id int) (*der.Album, error) {
	return s.repo.Get(id)
}

// SearchAlbums Search Albums
func (s *Service) SearchAlbums(query string) ([]*der.Album, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListAlbums List Albums
func (s *Service) ListAlbums() ([]*der.Album, error) {
	return s.repo.List()
}

// DeleteAlbum Delete Album
func (s *Service) DeleteAlbum(id int) error {
	u, err := s.GetAlbum(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateAlbum Update Album
func (s *Service) UpdateAlbum(e *der.Album) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
