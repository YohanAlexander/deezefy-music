package entity

import (
	"github.com/go-playground/validator/v10"
)

// SalvaAlbum entidade SalvaAlbum
type SalvaAlbum struct {
	Album   int    `validate:"required,gte=1"`
	Ouvinte string `validate:"required,email"`
	Artista string `validate:"required,email"`
}

// NewSalvaAlbum cria um novo SalvaAlbum
func NewSalvaAlbum(album int, ouvinte, artista string) (*SalvaAlbum, error) {
	sa := &SalvaAlbum{
		Album:   album,
		Ouvinte: ouvinte,
		Artista: artista,
	}
	err := sa.Validate()
	if err != nil {
		return nil, err
	}
	return sa, nil
}

// Validate valida os dados do SalvaAlbum
func (sa *SalvaAlbum) Validate() error {
	vld := validator.New()
	if err := vld.Struct(sa); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
