package salvaalbum

import der "github.com/yohanalexander/deezefy-music/entity/salvaalbum"

// SalvaAlbum interface
type SalvaAlbum interface {
	Get(album int, ouvinte, artista string) (*der.SalvaAlbum, error)
	GetByAlbum(album int) (*der.SalvaAlbum, error)
	GetByOuvinte(ouvinte string) (*der.SalvaAlbum, error)
	Search(query string) ([]*der.SalvaAlbum, error)
	List() ([]*der.SalvaAlbum, error)
	Create(e *der.SalvaAlbum) (int, string, string, error)
	Update(e *der.SalvaAlbum) error
	Delete(album int, ouvinte, artista string) error
}

// Repository interface
type Repository interface {
	SalvaAlbum
}

// UseCase interface
type UseCase interface {
	GetSalvaAlbum(album int, ouvinte, artista string) (*der.SalvaAlbum, error)
	GetSalvaAlbumByalbum(album int) (*der.SalvaAlbum, error)
	GetSalvaAlbumByOuvinte(ouvinte string) (*der.SalvaAlbum, error)
	SearchSalvaAlbums(query string) ([]*der.SalvaAlbum, error)
	ListSalvaAlbums() ([]*der.SalvaAlbum, error)
	CreateSalvaAlbum(album int, ouvinte, artista string) (int, string, string, error)
	UpdateSalvaAlbum(e *der.SalvaAlbum) error
	DeleteSalvaAlbum(album int, ouvinte, artista string) error
}
