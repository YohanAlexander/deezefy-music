package ouvinte

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

// CreateOuvinte Create Ouvinte
func (s *Service) CreateOuvinte(email, password, birthday, primeironome, sobrenome string) (string, error) {
	e, err := entity.NewOuvinte(email, password, birthday, primeironome, sobrenome)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetOuvinte Get Ouvinte
func (s *Service) GetOuvinte(email string) (*entity.Ouvinte, error) {
	return s.repo.Get(email)
}

// SearchOuvintes Search Ouvintes
func (s *Service) SearchOuvintes(query string) ([]*entity.Ouvinte, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListOuvintes List Ouvintes
func (s *Service) ListOuvintes() ([]*entity.Ouvinte, error) {
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
func (s *Service) UpdateOuvinte(e *entity.Ouvinte) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
