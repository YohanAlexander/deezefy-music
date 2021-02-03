package evento

import "github.com/yohanalexander/deezefy-music/entity"

// Evento interface
type Evento interface {
	Get(id int) (*entity.Evento, error)
	Search(query string) ([]*entity.Evento, error)
	List() ([]*entity.Evento, error)
	Create(e *entity.Evento) (int, error)
	Update(e *entity.Evento) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Evento
}

// UseCase interface
type UseCase interface {
	GetEvento(id int) (*entity.Evento, error)
	SearchEventos(query string) ([]*entity.Evento, error)
	ListEventos() ([]*entity.Evento, error)
	CreateEvento(usuario, nome string, id int) (int, error)
	UpdateEvento(e *entity.Evento) error
	DeleteEvento(id int) error
}
