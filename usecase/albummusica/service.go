package albummusica

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/albummusica"
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

// CreateAlbumMusica Create AlbumMusica
func (s *Service) CreateAlbumMusica(album, musica int, artista string) (int, int, string, error) {
	e, err := der.NewAlbumMusica(artista, album, musica)
	if err != nil {
		return 1, 1, err.Error(), err
	}
	return s.repo.Create(e)
}

// GetAlbumMusica Get AlbumMusica
func (s *Service) GetAlbumMusica(album, musica int, artista string) (*der.AlbumMusica, error) {
	return s.repo.Get(album, musica, artista)
}

// GetAlbumMusicaByAlbum Get AlbumMusica By Album
func (s *Service) GetAlbumMusicaByAlbum(album int) (*der.AlbumMusica, error) {
	return s.repo.GetByAlbum(album)
}

// GetAlbumMusicaByMusica Get AlbumMusica By Musica
func (s *Service) GetAlbumMusicaByMusica(Musica int) (*der.AlbumMusica, error) {
	return s.repo.GetByMusica(Musica)
}

// SearchAlbumMusicas Search AlbumMusicas
func (s *Service) SearchAlbumMusicas(query string) ([]*der.AlbumMusica, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListAlbumMusicas List AlbumMusicas
func (s *Service) ListAlbumMusicas() ([]*der.AlbumMusica, error) {
	return s.repo.List()
}

// DeleteAlbumMusica Delete AlbumMusica
func (s *Service) DeleteAlbumMusica(album, musica int, artista string) error {
	u, err := s.GetAlbumMusica(album, musica, artista)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(album, musica, artista)
}

// UpdateAlbumMusica Update AlbumMusica
func (s *Service) UpdateAlbumMusica(e *der.AlbumMusica) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
