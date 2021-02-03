package entity

import (
	"github.com/go-playground/validator/v10"
)

// Genero entidade Genero
type Genero struct {
	Nome   string `validate:"required,gte=1"`
	Estilo string `validate:"required,oneof=blues rock mpb samba sertanejo jazz classica"`
}

// NewGenero cria um novo Genero
func NewGenero(nome, estilo string) (*Genero, error) {
	g := &Genero{
		Nome:   nome,
		Estilo: estilo,
	}
	err := g.Validate()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Validate valida os dados do Genero
func (g *Genero) Validate() error {
	vld := validator.New()
	if err := vld.Struct(g); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
