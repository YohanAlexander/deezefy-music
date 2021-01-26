package ouvinte

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// Telefone entidade multivalorada telefone
type Telefone struct {
	Ouvinte  string `validate:"required,email"`
	Telefone string `validate:"e164"`
}

// NewTelefone cria um novo Telefone
func NewTelefone(ouvinte, telefone string) (*Telefone, error) {
	t := &Telefone{
		Ouvinte:  ouvinte,
		Telefone: telefone,
	}
	err := t.Validate()
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Validate valida os dados do Telefone
func (t *Telefone) Validate() error {
	vld := validator.New()
	if err := vld.Struct(t); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
