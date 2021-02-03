package entity

import (
	"github.com/go-playground/validator/v10"
)

// Ocorre entidade Ocorre
type Ocorre struct {
	Data    string `validate:"required,datetime=2006-01-02"`
	Local   int    `validate:"required,gte=1"`
	Evento  int    `validate:"required,gte=1"`
	Artista string `validate:"required,email"`
	Usuario string `validate:"required,email"`
}

// NewOcorre cria um novo Ocorre
func NewOcorre(data, artista, usuario string, local, evento int) (*Ocorre, error) {
	oc := &Ocorre{
		Data:    data,
		Local:   local,
		Evento:  evento,
		Artista: artista,
		Usuario: usuario,
	}
	err := oc.Validate()
	if err != nil {
		return nil, err
	}
	return oc, nil
}

// Validate valida os dados do Ocorre
func (oc *Ocorre) Validate() error {
	vld := validator.New()
	if err := vld.Struct(oc); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
