package usuario

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(email string) (*entity.Usuario, error)
	Search(query string) ([]*entity.Usuario, error)
	List() ([]*entity.Usuario, error)
}

// Write interface
type Write interface {
	Create(e *entity.Usuario) (string, error)
	Update(e *entity.Usuario) error
	Delete(email string) error
}


// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetUsuario(email string) (*entity.Usuario, error)
	SearchUsuarios(query string) ([]*entity.Usuario, error)
	ListUsuarios() ([]*entity.Usuario, error)
	CreateUsuario(email, password, birthday string) (string, error)
	UpdateUsuario(e *entity.Usuario) error
	DeleteUsuario(email string) error
}
