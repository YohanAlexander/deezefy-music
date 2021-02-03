package entity

import (
	"github.com/go-playground/validator/v10"
)

// Perfil entidade Perfil
type Perfil struct {
	Ouvinte               Ouvinte `validate:"required"`
	ID                    int     `validate:"required,gte=1"`
	InformacoesRelevantes string  `validate:"required,gte=1"`
}

// NewPerfil cria um novo Perfil
func NewPerfil(email, password, birthday, primeironome, sobrenome, informacoesrelevantes string, id int) (*Perfil, error) {
	o, err := NewOuvinte(email, password, birthday, primeironome, sobrenome)
	if err != nil {
		return nil, err
	}
	p := &Perfil{
		ID:                    id,
		Ouvinte:               *o,
		InformacoesRelevantes: informacoesrelevantes,
	}
	err = p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate valida os dados do Perfil
func (p *Perfil) Validate() error {
	vld := validator.New()
	if err := vld.Struct(p); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
