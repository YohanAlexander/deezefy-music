package generosfavoritos

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/generosfavoritos"
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

// CreateGenerosFavoritos Create GenerosFavoritos
func (s *Service) CreateGenerosFavoritos(perfil int, genero, ouvinte string) (int, string, string, error) {
	e, err := der.NewGenerosFavoritos(genero, ouvinte, perfil)
	if err != nil {
		return 1, err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetGenerosFavoritos Get GenerosFavoritos
func (s *Service) GetGenerosFavoritos(perfil int, genero, ouvinte string) (*der.GenerosFavoritos, error) {
	return s.repo.Get(perfil, genero, ouvinte)
}

// GetGenerosFavoritosByPerfil Get GenerosFavoritos By Perfil
func (s *Service) GetGenerosFavoritosByPerfil(perfil int) (*der.GenerosFavoritos, error) {
	return s.repo.GetByPerfil(perfil)
}

// GetGenerosFavoritosByGenero Get GenerosFavoritos By Genero
func (s *Service) GetGenerosFavoritosByGenero(genero string) (*der.GenerosFavoritos, error) {
	return s.repo.GetByGenero(genero)
}

// SearchGenerosFavoritos Search GenerosFavoritos
func (s *Service) SearchGenerosFavoritos(query string) ([]*der.GenerosFavoritos, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListGenerosFavoritos List GenerosFavoritos
func (s *Service) ListGenerosFavoritos() ([]*der.GenerosFavoritos, error) {
	return s.repo.List()
}

// DeleteGenerosFavoritos Delete GenerosFavoritos
func (s *Service) DeleteGenerosFavoritos(perfil int, genero, ouvinte string) error {
	u, err := s.GetGenerosFavoritos(perfil, genero, ouvinte)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(perfil, genero, ouvinte)
}

// UpdateGenerosFavoritos Update GenerosFavoritos
func (s *Service) UpdateGenerosFavoritos(e *der.GenerosFavoritos) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
