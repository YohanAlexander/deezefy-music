package artistagenero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistagenero"
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

// CreateArtistaGenero Create ArtistaGenero
func (s *Service) CreateArtistaGenero(artista, genero string) (string, string, error) {
	e, err := der.NewArtistaGenero(artista, genero)
	if err != nil {
		return err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetArtistaGenero Get ArtistaGenero
func (s *Service) GetArtistaGenero(artista, genero string) (*der.ArtistaGenero, error) {
	return s.repo.Get(artista, genero)
}

// GetArtistaGeneroByArtista Get ArtistaGenero By Artista
func (s *Service) GetArtistaGeneroByArtista(artista string) (*der.ArtistaGenero, error) {
	return s.repo.GetByArtista(artista)
}

// GetArtistaGeneroByGenero Get ArtistaGenero By Genero
func (s *Service) GetArtistaGeneroByGenero(genero string) (*der.ArtistaGenero, error) {
	return s.repo.GetByGenero(genero)
}

// SearchArtistaGeneros Search ArtistaGeneros
func (s *Service) SearchArtistaGeneros(query string) ([]*der.ArtistaGenero, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListArtistaGeneros List ArtistaGeneros
func (s *Service) ListArtistaGeneros() ([]*der.ArtistaGenero, error) {
	return s.repo.List()
}

// DeleteArtistaGenero Delete ArtistaGenero
func (s *Service) DeleteArtistaGenero(artista, genero string) error {
	u, err := s.GetArtistaGenero(artista, genero)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(artista, genero)
}

// UpdateArtistaGenero Update ArtistaGenero
func (s *Service) UpdateArtistaGenero(e *der.ArtistaGenero) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
