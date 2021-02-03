package entity

import (
	"github.com/go-playground/validator/v10"
)

// Curte entidade Curte
type Curte struct {
	Musica  int    `validate:"required,gte=1"`
	Ouvinte string `validate:"required,email"`
}

// NewCurte cria um novo Curte
func NewCurte(musica int, ouvinte string) (*Curte, error) {
	c := &Curte{
		Musica:  musica,
		Ouvinte: ouvinte,
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Validate valida os dados do Curte
func (c *Curte) Validate() error {
	vld := validator.New()
	if err := vld.Struct(c); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
