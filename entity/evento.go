package entity

import (
	"github.com/go-playground/validator/v10"
)

// Evento entidade Evento
type Evento struct {
	Usuario Usuario `validate:"required"`
	Nome    string  `validate:"required,gte=1"`
	ID      int     `validate:"required,gte=1"`
}

// NewEvento cria um novo Evento
func NewEvento(email, password, birthday, nome string, id int) (*Evento, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	e := &Evento{
		Usuario: *u,
		Nome:    nome,
		ID:      id,
	}
	err = e.Validate()
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Validate valida os dados do Evento
func (e *Evento) Validate() error {
	vld := validator.New()
	if err := vld.Struct(e); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
