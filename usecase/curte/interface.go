package curte

import der "github.com/yohanalexander/deezefy-music/entity/curte"

// Curte interface
type Curte interface {
	Get(musica int, ouvinte string) (*der.Curte, error)
	GetByMusica(musica int) (*der.Curte, error)
	GetByOuvinte(ouvinte string) (*der.Curte, error)
	Search(query string) ([]*der.Curte, error)
	List() ([]*der.Curte, error)
	Create(e *der.Curte) (int, string, error)
	Update(e *der.Curte) error
	Delete(musica int, ouvinte string) error
}

// Repository interface
type Repository interface {
	Curte
}

// UseCase interface
type UseCase interface {
	GetCurte(musica int, ouvinte string) (*der.Curte, error)
	GetCurteByMusica(Musica int) (*der.Curte, error)
	GetCurteByOuvinte(ouvinte string) (*der.Curte, error)
	SearchCurtes(query string) ([]*der.Curte, error)
	ListCurtes() ([]*der.Curte, error)
	CreateCurte(musica int, ouvinte string) (int, string, error)
	UpdateCurte(e *der.Curte) error
	DeleteCurte(musica int, ouvinte string) error
}
