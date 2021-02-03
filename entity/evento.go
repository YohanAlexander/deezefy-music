package entity

import (
	"github.com/go-playground/validator/v10"
)

// Evento entidade Evento
type Evento struct {
	Usuario Usuario `validate:"required"`
	ID      int     `validate:"required,gte=1"`
	Nome    string  `validate:"required,gte=1"`
	Data    string  `validate:"required,datetime=2006-01-02"`
}

// NewEvento cria um novo Evento
func NewEvento(email, password, birthday, nome, data string, id int) (*Evento, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	e := &Evento{
		Usuario: *u,
		ID:      id,
		Nome:    nome,
		Data:    data,
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
