package entity

import (
	"github.com/go-playground/validator/v10"
)

// Segue entidade Segue
type Segue struct {
	Artista string `validate:"required,email"`
	Ouvinte string `validate:"required,email"`
}

// NewSegue cria um novo Segue
func NewSegue(artista, ouvinte string) (*Segue, error) {
	s := &Segue{
		Artista: artista,
		Ouvinte: ouvinte,
	}
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Validate valida os dados do Segue
func (s *Segue) Validate() error {
	vld := validator.New()
	if err := vld.Struct(s); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
