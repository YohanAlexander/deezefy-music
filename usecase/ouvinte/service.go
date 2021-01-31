package ouvinte

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ouvinte"
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

// CreateOuvinte Create Ouvinte
func (s *Service) CreateOuvinte(usuario, primeironome, sobrenome string) (string, error) {
	e, err := der.NewOuvinte(usuario, primeironome, sobrenome)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetOuvinte Get Ouvinte
func (s *Service) GetOuvinte(email string) (*der.Ouvinte, error) {
	return s.repo.Get(email)
}

// SearchOuvintes Search Ouvintes
func (s *Service) SearchOuvintes(query string) ([]*der.Ouvinte, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListOuvintes List Ouvintes
func (s *Service) ListOuvintes() ([]*der.Ouvinte, error) {
	return s.repo.List()
}

// DeleteOuvinte Delete Ouvinte
func (s *Service) DeleteOuvinte(email string) error {
	u, err := s.GetOuvinte(email)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(email)
}

// UpdateOuvinte Update Ouvinte
func (s *Service) UpdateOuvinte(e *der.Ouvinte) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
