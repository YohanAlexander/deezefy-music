package musicagenero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicagenero"
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

// CreateMusicaGenero Create MusicaGenero
func (s *Service) CreateMusicaGenero(musica int, genero string) (int, string, error) {
	e, err := der.NewMusicaGenero(genero, musica)
	if err != nil {
		return 1, err.Error(), err
	}
	return s.repo.Create(e)
}

// GetMusicaGenero Get MusicaGenero
func (s *Service) GetMusicaGenero(musica int, ouvinte string) (*der.MusicaGenero, error) {
	return s.repo.Get(musica, ouvinte)
}

// GetMusicaGeneroByMusica Get MusicaGenero By Musica
func (s *Service) GetMusicaGeneroByMusica(musica int) (*der.MusicaGenero, error) {
	return s.repo.GetByMusica(musica)
}

// GetMusicaGeneroByGenero Get MusicaGenero By Genero
func (s *Service) GetMusicaGeneroByGenero(genero string) (*der.MusicaGenero, error) {
	return s.repo.GetByGenero(genero)
}

// SearchMusicaGeneros Search MusicaGeneros
func (s *Service) SearchMusicaGeneros(query string) ([]*der.MusicaGenero, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListMusicaGeneros List MusicaGeneros
func (s *Service) ListMusicaGeneros() ([]*der.MusicaGenero, error) {
	return s.repo.List()
}

// DeleteMusicaGenero Delete MusicaGenero
func (s *Service) DeleteMusicaGenero(musica int, ouvinte string) error {
	u, err := s.GetMusicaGenero(musica, ouvinte)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(musica, ouvinte)
}

// UpdateMusicaGenero Update MusicaGenero
func (s *Service) UpdateMusicaGenero(e *der.MusicaGenero) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
