package musica

import "github.com/yohanalexander/deezefy-music/entity"

// Musica interface
type Musica interface {
	Get(id int) (*entity.Musica, error)
	Search(query string) ([]*entity.Musica, error)
	List() ([]*entity.Musica, error)
	Create(e *entity.Musica) (int, error)
	Update(e *entity.Musica) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Musica
}

// UseCase interface
type UseCase interface {
	GetMusica(id int) (*entity.Musica, error)
	SearchMusicas(query string) ([]*entity.Musica, error)
	ListMusicas() ([]*entity.Musica, error)
	CreateMusica(id, duracao int, nome string) (int, error)
	UpdateMusica(e *entity.Musica) error
	DeleteMusica(id int) error
}
