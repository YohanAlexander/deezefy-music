package usuario

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

// CreateUsuario Create Usuario
func (s *Service) CreateUsuario(email, password, birthday string) (string, error) {
	e, err := entity.NewUsuario(email, password, birthday)
	if err != nil {
		return err.Error(), err
	}
	return s.repo.Create(e)
}

// GetUsuario Get Usuario
func (s *Service) GetUsuario(email string) (*entity.Usuario, error) {
	return s.repo.Get(email)
}

// SearchUsuarios Search Usuarios
func (s *Service) SearchUsuarios(query string) ([]*entity.Usuario, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListUsuarios List Usuarios
func (s *Service) ListUsuarios() ([]*entity.Usuario, error) {
	return s.repo.List()
}

// DeleteUsuario Delete Usuario
func (s *Service) DeleteUsuario(email string) error {
	u, err := s.GetUsuario(email)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(email)
}

// UpdateUsuario Update Usuario
func (s *Service) UpdateUsuario(e *entity.Usuario) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
