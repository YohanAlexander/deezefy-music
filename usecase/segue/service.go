package segue

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/segue"
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

// CreateSegue Create Segue
func (s *Service) CreateSegue(artista, ouvinte string) (string, string, error) {
	e, err := der.NewSegue(artista, ouvinte)
	if err != nil {
		return err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetSegue Get Segue
func (s *Service) GetSegue(artista, ouvinte string) (*der.Segue, error) {
	return s.repo.Get(artista, ouvinte)
}

// GetSegueByArtista Get Segue By Artista
func (s *Service) GetSegueByArtista(artista string) (*der.Segue, error) {
	return s.repo.GetByArtista(artista)
}

// GetSegueByOuvinte Get Segue By Ouvinte
func (s *Service) GetSegueByOuvinte(ouvinte string) (*der.Segue, error) {
	return s.repo.GetByOuvinte(ouvinte)
}

// SearchSegues Search Segues
func (s *Service) SearchSegues(query string) ([]*der.Segue, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListSegues List Segues
func (s *Service) ListSegues() ([]*der.Segue, error) {
	return s.repo.List()
}

// DeleteSegue Delete Segue
func (s *Service) DeleteSegue(artista, ouvinte string) error {
	u, err := s.GetSegue(artista, ouvinte)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(artista, ouvinte)
}

// UpdateSegue Update Segue
func (s *Service) UpdateSegue(e *der.Segue) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
