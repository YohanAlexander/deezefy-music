package evento

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(id int) (*entity.Evento, error)
	Search(query string) ([]*entity.Evento, error)
	List() ([]*entity.Evento, error)
}

// Write interface
type Write interface {
	Create(e *entity.Evento) (int, error)
	Update(e *entity.Evento) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetEvento(id int) (*entity.Evento, error)
	SearchEventos(query string) ([]*entity.Evento, error)
	ListEventos() ([]*entity.Evento, error)
	CreateEvento(email, password, birthday, nome, data, cidade, pais string, idlocal, idevento int) (int, error)
	UpdateEvento(e *entity.Evento) error
	DeleteEvento(id int) error
}
