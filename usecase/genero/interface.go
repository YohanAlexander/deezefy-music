package genero

import der "github.com/yohanalexander/deezefy-music/entity/genero"

// Genero interface
type Genero interface {
	Get(nome string) (*der.Genero, error)
	Search(query string) ([]*der.Genero, error)
	List() ([]*der.Genero, error)
	Create(e *der.Genero) (string, error)
	Update(e *der.Genero) error
	Delete(nome string) error
}

// Repository interface
type Repository interface {
	Genero
}

// UseCase interface
type UseCase interface {
	GetGenero(nome string) (*der.Genero, error)
	SearchGeneros(query string) ([]*der.Genero, error)
	ListGeneros() ([]*der.Genero, error)
	CreateGenero(nome, estilo string) (string, error)
	UpdateGenero(e *der.Genero) error
	DeleteGenero(nome string) error
}
