package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Evento entidade Evento
type Evento struct {
	Usuario Usuario `json:"usuario"`
	Local   Local   `json:"local"`
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Data    string  `json:"data"`
}

// AppendEvento adiciona presenter na lista
func AppendEvento(evento entity.Evento, eventos []Evento) []Evento {
	e := &Evento{}
	e.MakeEvento(evento)
	eventos = append(eventos, *e)
	return eventos
}

// MakeEvento seta os valores a partir da entidade
func (e *Evento) MakeEvento(evento entity.Evento) {
	e.Usuario.MakeUsuario(evento.Usuario)
	e.Local.MakeLocal(evento.Local)
	e.ID = evento.ID
	e.Nome = evento.Nome
	e.Data = evento.Data
}
