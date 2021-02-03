package entity

import (
	"github.com/go-playground/validator/v10"
)

// Perfil entidade Perfil
type Perfil struct {
	Ouvinte               string `validate:"required,email"`
	ID                    int    `validate:"required,gte=1"`
	InformacoesRelevantes string `validate:"required,gte=1"`
}

// NewPerfil cria um novo Perfil
func NewPerfil(ouvinte, informacoesrelevantes string, id int) (*Perfil, error) {
	p := &Perfil{
		ID:                    id,
		Ouvinte:               ouvinte,
		InformacoesRelevantes: informacoesrelevantes,
	}
	err := p.Validate()
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
