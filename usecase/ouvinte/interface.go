package ouvinte

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(email string) (*entity.Ouvinte, error)
	Search(query string) ([]*entity.Ouvinte, error)
	List() ([]*entity.Ouvinte, error)
}

// Write interface
type Write interface {
	Create(e *entity.Ouvinte) (string, error)
	Update(e *entity.Ouvinte) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetOuvinte(email string) (*entity.Ouvinte, error)
	SearchOuvintes(query string) ([]*entity.Ouvinte, error)
	ListOuvintes() ([]*entity.Ouvinte, error)
	CreateOuvinte(email, password, birthday, primeironome, sobrenome string) (string, error)
	UpdateOuvinte(e *entity.Ouvinte) error
	DeleteOuvinte(email string) error
}
