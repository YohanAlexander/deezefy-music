package criaplaylist

import der "github.com/yohanalexander/deezefy-music/entity/criaplaylist"

// CriaPlaylist interface
type CriaPlaylist interface {
	Get(playlist, usuario string) (*der.CriaPlaylist, error)
	GetByPlaylist(playlist string) (*der.CriaPlaylist, error)
	GetByUsuario(usuario string) (*der.CriaPlaylist, error)
	Search(query string) ([]*der.CriaPlaylist, error)
	List() ([]*der.CriaPlaylist, error)
	Create(e *der.CriaPlaylist) (string, string, error)
	Update(e *der.CriaPlaylist) error
	Delete(playlist, usuario string) error
}

// Repository interface
type Repository interface {
	CriaPlaylist
}

// UseCase interface
type UseCase interface {
	GetCriaPlaylist(playlist, usuario string) (*der.CriaPlaylist, error)
	GetCriaPlaylistByPlaylist(playlist string) (*der.CriaPlaylist, error)
	GetCriaPlaylistByUsuario(usuario string) (*der.CriaPlaylist, error)
	SearchCriaPlaylists(query string) ([]*der.CriaPlaylist, error)
	ListCriaPlaylists() ([]*der.CriaPlaylist, error)
	CreateCriaPlaylist(datacriacao, playlist, usuario string) (string, string, error)
	UpdateCriaPlaylist(e *der.CriaPlaylist) error
	DeleteCriaPlaylist(playlist, usuario string) error
}
