package curte

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/curte"
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

// CreateCurte Create Curte
func (s *Service) CreateCurte(musica int, ouvinte string) (int, string, error) {
	e, err := der.NewCurte(musica, ouvinte)
	if err != nil {
		return 1, err.Error(), err
	}
	return s.repo.Create(e)
}

// GetCurte Get Curte
func (s *Service) GetCurte(musica int, ouvinte string) (*der.Curte, error) {
	return s.repo.Get(musica, ouvinte)
}

// GetCurteByMusica Get Curte By Musica
func (s *Service) GetCurteByMusica(musica int) (*der.Curte, error) {
	return s.repo.GetByMusica(musica)
}

// GetCurteByOuvinte Get Curte By Ouvinte
func (s *Service) GetCurteByOuvinte(ouvinte string) (*der.Curte, error) {
	return s.repo.GetByOuvinte(ouvinte)
}

// SearchCurtes Search Curtes
func (s *Service) SearchCurtes(query string) ([]*der.Curte, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListCurtes List Curtes
func (s *Service) ListCurtes() ([]*der.Curte, error) {
	return s.repo.List()
}

// DeleteCurte Delete Curte
func (s *Service) DeleteCurte(musica int, ouvinte string) error {
	u, err := s.GetCurte(musica, ouvinte)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(musica, ouvinte)
}

// UpdateCurte Update Curte
func (s *Service) UpdateCurte(e *der.Curte) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
