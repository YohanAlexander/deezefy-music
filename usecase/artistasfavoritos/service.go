package artistasfavoritos

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistasfavoritos"
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

// CreateArtistasFavoritos Create ArtistasFavoritos
func (s *Service) CreateArtistasFavoritos(perfil int, ouvinte, artista string) (int, string, string, error) {
	e, err := der.NewArtistasFavoritos(perfil, ouvinte, artista)
	if err != nil {
		return 1, err.Error(), err.Error(), err
	}
	return s.repo.Create(e)
}

// GetArtistasFavoritos Get ArtistasFavoritos
func (s *Service) GetArtistasFavoritos(perfil int, ouvinte, artista string) (*der.ArtistasFavoritos, error) {
	return s.repo.Get(perfil, ouvinte, artista)
}

// GetArtistasFavoritosByPerfil Get ArtistasFavoritos By Perfil
func (s *Service) GetArtistasFavoritosByPerfil(perfil int) (*der.ArtistasFavoritos, error) {
	return s.repo.GetByPerfil(perfil)
}

// GetArtistasFavoritosByOuvinte Get ArtistasFavoritos By Ouvinte
func (s *Service) GetArtistasFavoritosByOuvinte(ouvinte string) (*der.ArtistasFavoritos, error) {
	return s.repo.GetByOuvinte(ouvinte)
}

// SearchArtistasFavoritoss Search ArtistasFavoritoss
func (s *Service) SearchArtistasFavoritoss(query string) ([]*der.ArtistasFavoritos, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListArtistasFavoritoss List ArtistasFavoritoss
func (s *Service) ListArtistasFavoritoss() ([]*der.ArtistasFavoritos, error) {
	return s.repo.List()
}

// DeleteArtistasFavoritos Delete ArtistasFavoritos
func (s *Service) DeleteArtistasFavoritos(perfil int, ouvinte, artista string) error {
	u, err := s.GetArtistasFavoritos(perfil, ouvinte, artista)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(perfil, ouvinte, artista)
}

// UpdateArtistasFavoritos Update ArtistasFavoritos
func (s *Service) UpdateArtistasFavoritos(e *der.ArtistasFavoritos) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
