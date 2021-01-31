package evento

import der "github.com/yohanalexander/deezefy-music/entity/evento"

// Evento interface
type Evento interface {
	Get(id int) (*der.Evento, error)
	Search(query string) ([]*der.Evento, error)
	List() ([]*der.Evento, error)
	Create(e *der.Evento) (int, error)
	Update(e *der.Evento) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Evento
}

// UseCase interface
type UseCase interface {
	GetEvento(id int) (*der.Evento, error)
	SearchEventos(query string) ([]*der.Evento, error)
	ListEventos() ([]*der.Evento, error)
	CreateEvento(usuario, nome string, id int) (int, error)
	UpdateEvento(e *der.Evento) error
	DeleteEvento(id int) error
}
