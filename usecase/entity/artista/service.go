package artista

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

// CreateArtista Create Artista
func (s *Service) CreateArtista(email, password, birthday, nomeartistico, biografia string, anoformacao int) (string, error) {
	e, err := entity.NewArtista(email, password, birthday, nomeartistico, biografia, anoformacao)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetArtista Get Artista
func (s *Service) GetArtista(email string) (*entity.Artista, error) {
	return s.repo.Get(email)
}

// SearchArtistas Search Artistas
func (s *Service) SearchArtistas(query string) ([]*entity.Artista, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListArtistas List Artistas
func (s *Service) ListArtistas() ([]*entity.Artista, error) {
	return s.repo.List()
}

// DeleteArtista Delete Artista
func (s *Service) DeleteArtista(email string) error {
	u, err := s.GetArtista(email)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(email)
}

// UpdateArtista Update Artista
func (s *Service) UpdateArtista(e *entity.Artista) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
