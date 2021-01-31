package genero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/genero"
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

// CreateGenero Create Genero
func (s *Service) CreateGenero(nome, estilo string) (string, error) {
	e, err := der.NewGenero(nome, estilo)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetGenero Get Genero
func (s *Service) GetGenero(email string) (*der.Genero, error) {
	return s.repo.Get(email)
}

// SearchGeneros Search Generos
func (s *Service) SearchGeneros(query string) ([]*der.Genero, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListGeneros List Generos
func (s *Service) ListGeneros() ([]*der.Genero, error) {
	return s.repo.List()
}

// DeleteGenero Delete Genero
func (s *Service) DeleteGenero(email string) error {
	u, err := s.GetGenero(email)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(email)
}

// UpdateGenero Update Genero
func (s *Service) UpdateGenero(e *der.Genero) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
