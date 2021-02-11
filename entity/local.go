package entity

import (
	"github.com/go-playground/validator/v10"
)

// Local entidade Local
type Local struct {
	ID     int    `json:"id" validate:"required,gte=1"`
	Cidade string `json:"cidade" validate:"required,gte=1"`
	Pais   string `json:"pais" validate:"required,gte=1"`
}

// NewLocal cria um novo Local
func NewLocal(cidade, pais string, id int) (*Local, error) {
	l := &Local{
		ID:     id,
		Cidade: cidade,
		Pais:   pais,
	}
	err := l.Validate()
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Validate valida os dados do Local
func (l *Local) Validate() error {
	vld := validator.New()
	if err := vld.Struct(l); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
