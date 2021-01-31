package ouvinte

import der "github.com/yohanalexander/deezefy-music/entity/ouvinte"

// Ouvinte interface
type Ouvinte interface {
	Get(email string) (*der.Ouvinte, error)
	Search(query string) ([]*der.Ouvinte, error)
	List() ([]*der.Ouvinte, error)
	Create(e *der.Ouvinte) (string, error)
	Update(e *der.Ouvinte) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Ouvinte
}

// UseCase interface
type UseCase interface {
	GetOuvinte(email string) (*der.Ouvinte, error)
	SearchOuvintes(query string) ([]*der.Ouvinte, error)
	ListOuvintes() ([]*der.Ouvinte, error)
	CreateOuvinte(usuario, primeironome, sobrenome string) (string, error)
	UpdateOuvinte(e *der.Ouvinte) error
	DeleteOuvinte(email string) error
}
