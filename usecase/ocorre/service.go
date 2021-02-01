package ocorre

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ocorre"
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

// CreateOcorre Create Ocorre
func (s *Service) CreateOcorre(data, artista, usuario string, local, evento int) (string, string, int, int, error) {
	e, err := der.NewOcorre(data, artista, usuario, local, evento)
	if err != nil {
		return err.Error(), err.Error(), 1, 1, err
	}
	return s.repo.Create(e)
}

// GetOcorre Get Ocorre
func (s *Service) GetOcorre(artista, usuario string, local, evento int) (*der.Ocorre, error) {
	return s.repo.Get(artista, usuario, local, evento)
}

// GetOcorreByLocal Get Ocorre By Local
func (s *Service) GetOcorreByLocal(local int) (*der.Ocorre, error) {
	return s.repo.GetByLocal(local)
}

// GetOcorreByEvento Get Ocorre By Evento
func (s *Service) GetOcorreByEvento(evento int) (*der.Ocorre, error) {
	return s.repo.GetByEvento(evento)
}

// GetOcorreByArtista Get Ocorre By Artista
func (s *Service) GetOcorreByArtista(artista string) (*der.Ocorre, error) {
	return s.repo.GetByArtista(artista)
}

// SearchOcorres Search Ocorres
func (s *Service) SearchOcorres(query string) ([]*der.Ocorre, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListOcorres List Ocorres
func (s *Service) ListOcorres() ([]*der.Ocorre, error) {
	return s.repo.List()
}

// DeleteOcorre Delete Ocorre
func (s *Service) DeleteOcorre(artista, usuario string, local, evento int) error {
	u, err := s.GetOcorre(artista, usuario, local, evento)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(artista, usuario, local, evento)
}

// UpdateOcorre Update Ocorre
func (s *Service) UpdateOcorre(e *der.Ocorre) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
