package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Evento entidade Evento
type Evento struct {
	Usuario entity.Usuario `json:"usuario"`
	Local   entity.Local   `json:"local"`
	ID      int            `json:"id"`
	Nome    string         `json:"nome"`
	Data    string         `json:"data"`
}

// AppendEvento adiciona presenter na lista
func AppendEvento(evento entity.Evento, eventos []*Evento) []*Evento {
	e := &Evento{}
	e.GetEvento(evento)
	eventos = append(eventos, e)
	return eventos
}

// GetEvento seta os valores a partir da entidade
func (e *Evento) GetEvento(evento entity.Evento) {
	e.Usuario = evento.Usuario
	e.Local = evento.Local
	e.ID = evento.ID
	e.Nome = evento.Nome
	e.Data = evento.Data
}
