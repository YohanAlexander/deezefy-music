package local

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(id int) (*entity.Local, error)
	Search(query string) ([]*entity.Local, error)
	List() ([]*entity.Local, error)
}

// Write interface
type Write interface {
	Create(e *entity.Local) (int, error)
	Update(e *entity.Local) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetLocal(id int) (*entity.Local, error)
	SearchLocals(query string) ([]*entity.Local, error)
	ListLocals() ([]*entity.Local, error)
	CreateLocal(cidade, pais string, id int) (int, error)
	UpdateLocal(e *entity.Local) error
	DeleteLocal(id int) error
}
