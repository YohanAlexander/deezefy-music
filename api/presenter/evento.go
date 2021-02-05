package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Evento entidade Evento
type Evento struct {
	Usuario entity.Usuario `json:"usuario"`
	ID      int            `json:"id"`
	Nome    string         `json:"nome"`
	Data    string         `json:"data"`
}

// GetEvento seta os valores a partir da entidade
func (e *Evento) GetEvento(evento entity.Evento) {
	e.Usuario = evento.Usuario
	e.ID = evento.ID
	e.Nome = evento.Nome
	e.Data = evento.Data
}
