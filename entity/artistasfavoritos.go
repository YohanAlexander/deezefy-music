package entity

import (
	"github.com/go-playground/validator/v10"
)

// ArtistasFavoritos entidade ArtistasFavoritos
type ArtistasFavoritos struct {
	Perfil  int    `validate:"required,gte=1"`
	Ouvinte string `validate:"required,email"`
	Artista string `validate:"required,email"`
}

// NewArtistasFavoritos cria um novo ArtistasFavoritos
func NewArtistasFavoritos(perfil int, ouvinte, artista string) (*ArtistasFavoritos, error) {
	f := &ArtistasFavoritos{
		Perfil:  perfil,
		Ouvinte: ouvinte,
		Artista: artista,
	}
	err := f.Validate()
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Validate valida os dados do ArtistasFavoritos
func (f *ArtistasFavoritos) Validate() error {
	vld := validator.New()
	if err := vld.Struct(f); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
