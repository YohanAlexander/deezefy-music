package salvaplaylist

import der "github.com/yohanalexander/deezefy-music/entity/salvaplaylist"

// SalvaPlaylist interface
type SalvaPlaylist interface {
	Get(playlist, ouvinte string) (*der.SalvaPlaylist, error)
	GetByPlaylist(playlist string) (*der.SalvaPlaylist, error)
	GetByOuvinte(ouvinte string) (*der.SalvaPlaylist, error)
	Search(query string) ([]*der.SalvaPlaylist, error)
	List() ([]*der.SalvaPlaylist, error)
	Create(e *der.SalvaPlaylist) (string, string, error)
	Update(e *der.SalvaPlaylist) error
	Delete(playlist, ouvinte string) error
}

// Repository interface
type Repository interface {
	SalvaPlaylist
}

// UseCase interface
type UseCase interface {
	GetSalvaPlaylist(playlist, ouvinte string) (*der.SalvaPlaylist, error)
	GetSalvaPlaylistByPlaylist(playlist string) (*der.SalvaPlaylist, error)
	GetSalvaPlaylistByOuvinte(ouvinte string) (*der.SalvaPlaylist, error)
	SearchSalvaPlaylists(query string) ([]*der.SalvaPlaylist, error)
	ListSalvaPlaylists() ([]*der.SalvaPlaylist, error)
	CreateSalvaPlaylist(playlist, ouvinte string) (string, string, error)
	UpdateSalvaPlaylist(e *der.SalvaPlaylist) error
	DeleteSalvaPlaylist(playlist, ouvinte string) error
}
