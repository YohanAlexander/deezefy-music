package artista

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(email string) (*entity.Artista, error)
	Search(query string) ([]*entity.Artista, error)
	List() ([]*entity.Artista, error)
}

// Write interface
type Write interface {
	Create(e *entity.Artista) (string, error)
	Update(e *entity.Artista) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetArtista(email string) (*entity.Artista, error)
	SearchArtistas(query string) ([]*entity.Artista, error)
	ListArtistas() ([]*entity.Artista, error)
	CreateArtista(email, password, birthday, nomeartistico, biografia string, anoformacao int) (string, error)
	UpdateArtista(e *entity.Artista) error
	DeleteArtista(email string) error
}
