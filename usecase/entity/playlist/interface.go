package playlist

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(nome string) (*entity.Playlist, error)
	Search(query string) ([]*entity.Playlist, error)
	List() ([]*entity.Playlist, error)
}

// Write interface
type Write interface {
	Create(e *entity.Playlist) (string, error)
	Update(e *entity.Playlist) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetPlaylist(nome string) (*entity.Playlist, error)
	SearchPlaylists(query string) ([]*entity.Playlist, error)
	ListPlaylists() ([]*entity.Playlist, error)
	CreatePlaylist(email, password, birthday, nome, status, datacriacao string) (string, error)
	UpdatePlaylist(e *entity.Playlist) error
	DeletePlaylist(nome string) error
}
