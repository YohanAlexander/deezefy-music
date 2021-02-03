package entity

import (
	"github.com/go-playground/validator/v10"
)

// Artista entidade Artista
type Artista struct {
	Usuario       string `validate:"required,email"`
	NomeArtistico string `validate:"required,min=2"`
	Biografia     string `validate:"gte=10"`
	AnoFormacao   int    `validate:"gte=1900"`
}

// NewArtista cria um novo Artista
func NewArtista(usuario, nomeartistico, biografia string, anoformacao int) (*Artista, error) {
	a := &Artista{
		Usuario:       usuario,
		NomeArtistico: nomeartistico,
		Biografia:     biografia,
		AnoFormacao:   anoformacao,
	}
	err := a.Validate()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Validate valida os dados do Artista
func (a *Artista) Validate() error {
	vld := validator.New()
	if err := vld.Struct(a); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
