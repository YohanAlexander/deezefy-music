package musica

import der "github.com/yohanalexander/deezefy-music/entity/musica"

// Musica interface
type Musica interface {
	Get(id int) (*der.Musica, error)
	Search(query string) ([]*der.Musica, error)
	List() ([]*der.Musica, error)
	Create(e *der.Musica) (int, error)
	Update(e *der.Musica) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Musica
}

// UseCase interface
type UseCase interface {
	GetMusica(id int) (*der.Musica, error)
	SearchMusicas(query string) ([]*der.Musica, error)
	ListMusicas() ([]*der.Musica, error)
	CreateMusica(id, duracao int, nome string) (int, error)
	UpdateMusica(e *der.Musica) error
	DeleteMusica(id int) error
}
