package local

import der "github.com/yohanalexander/deezefy-music/entity/local"

// Local interface
type Local interface {
	Get(id int) (*der.Local, error)
	Search(query string) ([]*der.Local, error)
	List() ([]*der.Local, error)
	Create(e *der.Local) (int, error)
	Update(e *der.Local) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Local
}

// UseCase interface
type UseCase interface {
	GetLocal(id int) (*der.Local, error)
	SearchLocals(query string) ([]*der.Local, error)
	ListLocals() ([]*der.Local, error)
	CreateLocal(cidade, pais string, id int) (int, error)
	UpdateLocal(e *der.Local) error
	DeleteLocal(id int) error
}
