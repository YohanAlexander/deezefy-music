package playlist

import "github.com/yohanalexander/deezefy-music/entity"

// Playlist interface
type Playlist interface {
	Get(nome string) (*entity.Playlist, error)
	Search(query string) ([]*entity.Playlist, error)
	List() ([]*entity.Playlist, error)
	Create(e *entity.Playlist) (string, error)
	Update(e *entity.Playlist) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Playlist
}

// UseCase interface
type UseCase interface {
	GetPlaylist(nome string) (*entity.Playlist, error)
	SearchPlaylists(query string) ([]*entity.Playlist, error)
	ListPlaylists() ([]*entity.Playlist, error)
	CreatePlaylist(nome, status string) (string, error)
	UpdatePlaylist(e *entity.Playlist) error
	DeletePlaylist(nome string) error
}
