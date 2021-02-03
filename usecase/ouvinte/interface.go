package ouvinte

import "github.com/yohanalexander/deezefy-music/entity"

// Ouvinte interface
type Ouvinte interface {
	Get(email string) (*entity.Ouvinte, error)
	Search(query string) ([]*entity.Ouvinte, error)
	List() ([]*entity.Ouvinte, error)
	Create(e *entity.Ouvinte) (string, error)
	Update(e *entity.Ouvinte) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Ouvinte
}

// UseCase interface
type UseCase interface {
	GetOuvinte(email string) (*entity.Ouvinte, error)
	SearchOuvintes(query string) ([]*entity.Ouvinte, error)
	ListOuvintes() ([]*entity.Ouvinte, error)
	CreateOuvinte(usuario, primeironome, sobrenome string) (string, error)
	UpdateOuvinte(e *entity.Ouvinte) error
	DeleteOuvinte(email string) error
}
