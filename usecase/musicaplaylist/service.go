package musicaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicaplaylist"
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

// CreateMusicaPlaylist Create MusicaPlaylist
func (s *Service) CreateMusicaPlaylist(musica int, playlist string) (int, string, error) {
	e, err := der.NewMusicaPlaylist(playlist, musica)
	if err != nil {
		return 1, err.Error(), err
	}
	return s.repo.Create(e)
}

// GetMusicaPlaylist Get MusicaPlaylist
func (s *Service) GetMusicaPlaylist(musica int, playlist string) (*der.MusicaPlaylist, error) {
	return s.repo.Get(musica, playlist)
}

// GetMusicaPlaylistByMusica Get MusicaPlaylist By Musica
func (s *Service) GetMusicaPlaylistByMusica(musica int) (*der.MusicaPlaylist, error) {
	return s.repo.GetByMusica(musica)
}

// GetMusicaPlaylistByPlaylist Get MusicaPlaylist By Playlist
func (s *Service) GetMusicaPlaylistByPlaylist(playlist string) (*der.MusicaPlaylist, error) {
	return s.repo.GetByPlaylist(playlist)
}

// SearchMusicaPlaylists Search MusicaPlaylists
func (s *Service) SearchMusicaPlaylists(query string) ([]*der.MusicaPlaylist, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListMusicaPlaylists List MusicaPlaylists
func (s *Service) ListMusicaPlaylists() ([]*der.MusicaPlaylist, error) {
	return s.repo.List()
}

// DeleteMusicaPlaylist Delete MusicaPlaylist
func (s *Service) DeleteMusicaPlaylist(musica int, playlist string) error {
	u, err := s.GetMusicaPlaylist(musica, playlist)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(musica, playlist)
}

// UpdateMusicaPlaylist Update MusicaPlaylist
func (s *Service) UpdateMusicaPlaylist(e *der.MusicaPlaylist) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(e)
}
