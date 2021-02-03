package genero

import "github.com/yohanalexander/deezefy-music/entity"

// Genero interface
type Genero interface {
	Get(nome string) (*entity.Genero, error)
	Search(query string) ([]*entity.Genero, error)
	List() ([]*entity.Genero, error)
	Create(e *entity.Genero) (string, error)
	Update(e *entity.Genero) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Genero
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
