package ocorre

import der "github.com/yohanalexander/deezefy-music/entity/ocorre"

// Ocorre interface
type Ocorre interface {
	Get(artista, usuario string, local, evento int) (*der.Ocorre, error)
	GetByLocal(local int) (*der.Ocorre, error)
	GetByEvento(evento int) (*der.Ocorre, error)
	GetByArtista(artista string) (*der.Ocorre, error)
	Search(query string) ([]*der.Ocorre, error)
	List() ([]*der.Ocorre, error)
	Create(e *der.Ocorre) (string, string, int, int, error)
	Update(e *der.Ocorre) error
	Delete(artista, usuario string, local, evento int) error
}

// Repository interface
type Repository interface {
	Ocorre
}

// UseCase interface
type UseCase interface {
	GetOcorre(artista, usuario string, local, evento int) (*der.Ocorre, error)
	GetOcorreByLocal(local int) (*der.Ocorre, error)
	GetOcorreByEvento(evento int) (*der.Ocorre, error)
	GetOcorreByArtista(artista string) (*der.Ocorre, error)
	SearchOcorres(query string) ([]*der.Ocorre, error)
	ListOcorres() ([]*der.Ocorre, error)
	CreateOcorre(data, artista, usuario string, local, evento int) (string, string, int, int, error)
	UpdateOcorre(e *der.Ocorre) error
	DeleteOcorre(artista, usuario string, local, evento int) error
}
