package entity

import (
	"github.com/go-playground/validator/v10"
)

// AlbumMusica entidade AlbumMusica
type AlbumMusica struct {
	Album   int    `validate:"required,gte=1"`
	Artista string `validate:"required,email"`
	Musica  int    `validate:"required,gte=1"`
}

// NewAlbumMusica cria um novo AlbumMusica
func NewAlbumMusica(artista string, album, musica int) (*AlbumMusica, error) {
	am := &AlbumMusica{
		Album:   album,
		Artista: artista,
		Musica:  musica,
	}
	err := am.Validate()
	if err != nil {
		return nil, err
	}
	return am, nil
}

// Validate valida os dados do AlbumMusica
func (am *AlbumMusica) Validate() error {
	vld := validator.New()
	if err := vld.Struct(am); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
