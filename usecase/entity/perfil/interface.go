package perfil

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(id int) (*entity.Perfil, error)
	Search(query string) ([]*entity.Perfil, error)
	List() ([]*entity.Perfil, error)
}

// Write interface
type Write interface {
	Create(e *entity.Perfil) (int, error)
	Update(e *entity.Perfil) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetPerfil(id int) (*entity.Perfil, error)
	SearchPerfils(query string) ([]*entity.Perfil, error)
	ListPerfils() ([]*entity.Perfil, error)
	CreatePerfil(email, password, birthday, primeironome, sobrenome, informacoesrelevantes string, id int) (int, error)
	UpdatePerfil(e *entity.Perfil) error
	DeletePerfil(id int) error
}
