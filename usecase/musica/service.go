package musica

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musica"
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

// CreateMusica Create Musica
func (s *Service) CreateMusica(id, duracao int, nome string) (int, error) {
	e, err := der.NewMusica(id, duracao, nome)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetMusica Get Musica
func (s *Service) GetMusica(id int) (*der.Musica, error) {
	return s.repo.Get(id)
}

// SearchMusicas Search Musicas
func (s *Service) SearchMusicas(query string) ([]*der.Musica, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListMusicas List Musicas
func (s *Service) ListMusicas() ([]*der.Musica, error) {
	return s.repo.List()
}

// DeleteMusica Delete Musica
func (s *Service) DeleteMusica(id int) error {
	u, err := s.GetMusica(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateMusica Update Musica
func (s *Service) UpdateMusica(e *der.Musica) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
