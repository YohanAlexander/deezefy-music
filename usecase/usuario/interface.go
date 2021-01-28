package usuario

import der "github.com/yohanalexander/deezefy-music/entity/usuario"

// Usuario interface
type Usuario interface {
	Get(email string) (*der.Usuario, error)
	Search(query string) ([]*der.Usuario, error)
	List() ([]*der.Usuario, error)
	Create(e *der.Usuario) (string, error)
	Update(e *der.Usuario) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Usuario
}

// UseCase interface
type UseCase interface {
	GetUsuario(email string) (*der.Usuario, error)
	SearchUsuarios(query string) ([]*der.Usuario, error)
	ListUsuarios() ([]*der.Usuario, error)
	CreateUsuario(email, password, birthday string) (string, error)
	UpdateUsuario(e *der.Usuario) error
	DeleteUsuario(email string) error
}
