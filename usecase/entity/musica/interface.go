package musica

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(id int) (*entity.Musica, error)
	Search(query string) ([]*entity.Musica, error)
	List() ([]*entity.Musica, error)
}

// Write interface
type Write interface {
	Create(e *entity.Musica) (int, error)
	Update(e *entity.Musica) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetMusica(id int) (*entity.Musica, error)
	SearchMusicas(query string) ([]*entity.Musica, error)
	ListMusicas() ([]*entity.Musica, error)
	CreateMusica(nome string, duracao, id int) (int, error)
	UpdateMusica(e *entity.Musica) error
	DeleteMusica(id int) error
}
