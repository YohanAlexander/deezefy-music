package evento

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/evento"
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

// CreateEvento Create Evento
func (s *Service) CreateEvento(usuario, nome string, id int) (int, error) {
	e, err := der.NewEvento(usuario, nome, id)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetEvento Get Evento
func (s *Service) GetEvento(id int) (*der.Evento, error) {
	return s.repo.Get(id)
}

// SearchEventos Search Eventos
func (s *Service) SearchEventos(query string) ([]*der.Evento, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListEventos List Eventos
func (s *Service) ListEventos() ([]*der.Evento, error) {
	return s.repo.List()
}

// DeleteEvento Delete Evento
func (s *Service) DeleteEvento(id int) error {
	u, err := s.GetEvento(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateEvento Update Evento
func (s *Service) UpdateEvento(e *der.Evento) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
