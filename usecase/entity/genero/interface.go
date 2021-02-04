package genero

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(nome string) (*entity.Genero, error)
	Search(query string) ([]*entity.Genero, error)
	List() ([]*entity.Genero, error)
}

// Write interface
type Write interface {
	Create(e *entity.Genero) (string, error)
	Update(e *entity.Genero) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetGenero(nome string) (*entity.Genero, error)
	SearchGeneros(query string) ([]*entity.Genero, error)
	ListGeneros() ([]*entity.Genero, error)
	CreateGenero(nome, estilo string) (string, error)
	UpdateGenero(e *entity.Genero) error
	DeleteGenero(nome string) error
}
