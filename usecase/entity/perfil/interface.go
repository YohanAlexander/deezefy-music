package perfil

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(email string) (*entity.Perfil, error)
	Search(query string) ([]*entity.Perfil, error)
	List() ([]*entity.Perfil, error)
}

// Write interface
type Write interface {
	Create(e *entity.Perfil) (string, error)
	Update(e *entity.Perfil) error
	Delete(email string) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetPerfil(email string) (*entity.Perfil, error)
	SearchPerfils(query string) ([]*entity.Perfil, error)
	ListPerfils() ([]*entity.Perfil, error)
	CreatePerfil(email, password, birthday, primeironome, sobrenome, informacoesrelevantes string, id int) (string, error)
	UpdatePerfil(e *entity.Perfil) error
	DeletePerfil(email string) error
}
