package grava

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/grava"
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

// CreateGrava Create Grava
func (s *Service) CreateGrava(musica int, artista string) (int, string, error) {
	e, err := der.NewGrava(musica, artista)
	if err != nil {
		return 1, err.Error(), err
	}
	return s.repo.Create(e)
}

// GetGrava Get Grava
func (s *Service) GetGrava(musica int, artista string) (*der.Grava, error) {
	return s.repo.Get(musica, artista)
}

// GetGravaByMusica Get Grava By Musica
func (s *Service) GetGravaByMusica(musica int) (*der.Grava, error) {
	return s.repo.GetByMusica(musica)
}

// GetGravaByArtista Get Grava By Artista
func (s *Service) GetGravaByArtista(artista string) (*der.Grava, error) {
	return s.repo.GetByArtista(artista)
}

// SearchGravas Search Gravas
func (s *Service) SearchGravas(query string) ([]*der.Grava, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListGravas List Gravas
func (s *Service) ListGravas() ([]*der.Grava, error) {
	return s.repo.List()
}

// DeleteGrava Delete Grava
func (s *Service) DeleteGrava(musica int, artista string) error {
	u, err := s.GetGrava(musica, artista)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(musica, artista)
}

// UpdateGrava Update Grava
func (s *Service) UpdateGrava(e *der.Grava) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
