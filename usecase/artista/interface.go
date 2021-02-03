package artista

import "github.com/yohanalexander/deezefy-music/entity"

// Artista interface
type Artista interface {
	Get(email string) (*entity.Artista, error)
	Search(query string) ([]*entity.Artista, error)
	List() ([]*entity.Artista, error)
	Create(e *entity.Artista) (string, error)
	Update(e *entity.Artista) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Artista
}

// UseCase interface
type UseCase interface {
	GetArtista(email string) (*entity.Artista, error)
	SearchArtistas(query string) ([]*entity.Artista, error)
	ListArtistas() ([]*entity.Artista, error)
	CreateArtista(usuario, nomeartistico, biografia string, anoformacao int) (string, error)
	UpdateArtista(e *entity.Artista) error
	DeleteArtista(email string) error
}
