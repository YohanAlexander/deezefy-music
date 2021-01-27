package musicagenero

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// MusicaGenero entidade MusicaGenero
type MusicaGenero struct {
	Musica int    `validate:"required,gte=1"`
	Genero string `validate:"required,gte=1"`
}

// NewMusicaGenero cria um novo MusicaGenero
func NewMusicaGenero(genero string, musica int) (*MusicaGenero, error) {
	mg := &MusicaGenero{
		Genero: genero,
		Musica: musica,
	}
	err := mg.Validate()
	if err != nil {
		return nil, err
	}
	return mg, nil
}

// Validate valida os dados do MusicaGenero
func (mg *MusicaGenero) Validate() error {
	vld := validator.New()
	if err := vld.Struct(mg); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
