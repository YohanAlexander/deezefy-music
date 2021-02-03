package entity

import (
	"github.com/go-playground/validator/v10"
)

// Artista entidade Artista
type Artista struct {
	Usuario       Usuario `validate:"required"`
	NomeArtistico string  `validate:"required,min=1"`
	Biografia     string  `validate:"gte=1"`
	AnoFormacao   int     `validate:"gte=1000"`
}

// NewArtista cria um novo Artista
func NewArtista(email, password, birthday, nomeartistico, biografia string, anoformacao int) (*Artista, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	a := &Artista{
		Usuario:       *u,
		NomeArtistico: nomeartistico,
		Biografia:     biografia,
		AnoFormacao:   anoformacao,
	}
	err = a.Validate()
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
