package perfil

import "github.com/yohanalexander/deezefy-music/entity"

// Perfil interface
type Perfil interface {
	Get(id int) (*entity.Perfil, error)
	Search(query string) ([]*entity.Perfil, error)
	List() ([]*entity.Perfil, error)
	Create(e *entity.Perfil) (int, error)
	Update(e *entity.Perfil) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Perfil
}

// UseCase interface
type UseCase interface {
	GetPerfil(id int) (*entity.Perfil, error)
	SearchPerfils(query string) ([]*entity.Perfil, error)
	ListPerfils() ([]*entity.Perfil, error)
	CreatePerfil(ouvinte, informacoesrelevantes string, id int) (int, error)
	UpdatePerfil(e *entity.Perfil) error
	DeletePerfil(id int) error
}
