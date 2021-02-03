package perfil

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

// CreatePerfil Create Perfil
func (s *Service) CreatePerfil(ouvinte, informacoesrelevantes string, id int) (int, error) {
	e, err := entity.NewPerfil(ouvinte, informacoesrelevantes, id)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetPerfil Get Perfil
func (s *Service) GetPerfil(id int) (*entity.Perfil, error) {
	return s.repo.Get(id)
}

// SearchPerfils Search Perfils
func (s *Service) SearchPerfils(query string) ([]*entity.Perfil, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListPerfils List Perfils
func (s *Service) ListPerfils() ([]*entity.Perfil, error) {
	return s.repo.List()
}

// DeletePerfil Delete Perfil
func (s *Service) DeletePerfil(id int) error {
	u, err := s.GetPerfil(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdatePerfil Update Perfil
func (s *Service) UpdatePerfil(e *entity.Perfil) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
