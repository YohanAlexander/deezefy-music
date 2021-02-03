package entity

import (
	"github.com/go-playground/validator/v10"
)

// ArtistaGenero entidade ArtistaGenero
type ArtistaGenero struct {
	Artista string `validate:"required,email"`
	Genero  string `validate:"required,gte=1"`
}

// NewArtistaGenero cria um novo ArtistaGenero
func NewArtistaGenero(artista, genero string) (*ArtistaGenero, error) {
	ag := &ArtistaGenero{
		Artista: artista,
		Genero:  genero,
	}
	err := ag.Validate()
	if err != nil {
		return nil, err
	}
	return ag, nil
}

// Validate valida os dados do ArtistaGenero
func (ag *ArtistaGenero) Validate() error {
	vld := validator.New()
	if err := vld.Struct(ag); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
