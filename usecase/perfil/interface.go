package perfil

import der "github.com/yohanalexander/deezefy-music/entity/perfil"

// Perfil interface
type Perfil interface {
	Get(id int) (*der.Perfil, error)
	Search(query string) ([]*der.Perfil, error)
	List() ([]*der.Perfil, error)
	Create(e *der.Perfil) (int, error)
	Update(e *der.Perfil) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Perfil
}

// UseCase interface
type UseCase interface {
	GetPerfil(id int) (*der.Perfil, error)
	SearchPerfils(query string) ([]*der.Perfil, error)
	ListPerfils() ([]*der.Perfil, error)
	CreatePerfil(ouvinte, informacoesrelevantes string, id int) (int, error)
	UpdatePerfil(e *der.Perfil) error
	DeletePerfil(id int) error
}
