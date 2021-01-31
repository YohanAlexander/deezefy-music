package artista

import der "github.com/yohanalexander/deezefy-music/entity/artista"

// Artista interface
type Artista interface {
	Get(email string) (*der.Artista, error)
	Search(query string) ([]*der.Artista, error)
	List() ([]*der.Artista, error)
	Create(e *der.Artista) (string, error)
	Update(e *der.Artista) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Artista
}

// UseCase interface
type UseCase interface {
	GetArtista(email string) (*der.Artista, error)
	SearchArtistas(query string) ([]*der.Artista, error)
	ListArtistas() ([]*der.Artista, error)
	CreateArtista(usuario, nomeartistico, biografia string, anoformacao int) (string, error)
	UpdateArtista(e *der.Artista) error
	DeleteArtista(email string) error
}
