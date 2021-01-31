package playlist

import der "github.com/yohanalexander/deezefy-music/entity/playlist"

// Playlist interface
type Playlist interface {
	Get(nome string) (*der.Playlist, error)
	Search(query string) ([]*der.Playlist, error)
	List() ([]*der.Playlist, error)
	Create(e *der.Playlist) (string, error)
	Update(e *der.Playlist) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Playlist
}

// UseCase interface
type UseCase interface {
	GetPlaylist(nome string) (*der.Playlist, error)
	SearchPlaylists(query string) ([]*der.Playlist, error)
	ListPlaylists() ([]*der.Playlist, error)
	CreatePlaylist(nome, status string) (string, error)
	UpdatePlaylist(e *der.Playlist) error
	DeletePlaylist(nome string) error
}
