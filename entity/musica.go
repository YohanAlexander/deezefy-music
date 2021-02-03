package entity

import (
	"github.com/go-playground/validator/v10"
)

// Musica entidade Musica
type Musica struct {
	ID      int    `validate:"required,gte=1"`
	Duracao int    `validate:"required,gte=100"`
	Nome    string `validate:"required,gte=1"`
}

// NewMusica cria um novo Musica
func NewMusica(id, duracao int, nome string) (*Musica, error) {
	m := &Musica{
		ID:      id,
		Duracao: duracao,
		Nome:    nome,
	}
	err := m.Validate()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Validate valida os dados do Musica
func (m *Musica) Validate() error {
	vld := validator.New()
	if err := vld.Struct(m); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
