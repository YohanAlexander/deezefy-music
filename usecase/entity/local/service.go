package local

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

// CreateLocal Create Local
func (s *Service) CreateLocal(cidade, pais string, id int) (int, error) {
	e, err := entity.NewLocal(cidade, pais, id)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetLocal Get Local
func (s *Service) GetLocal(id int) (*entity.Local, error) {
	return s.repo.Get(id)
}

// SearchLocals Search Locals
func (s *Service) SearchLocals(query string) ([]*entity.Local, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListLocals List Locals
func (s *Service) ListLocals() ([]*entity.Local, error) {
	return s.repo.List()
}

// DeleteLocal Delete Local
func (s *Service) DeleteLocal(id int) error {
	u, err := s.GetLocal(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateLocal Update Local
func (s *Service) UpdateLocal(e *entity.Local) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
