package generosfavoritos

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// GenerosFavoritos entidade GenerosFavoritos
type GenerosFavoritos struct {
	Genero  string `validate:"required,gte=1"`
	Perfil  int    `validate:"required,gte=1"`
	Ouvinte string `validate:"required,email"`
}

// NewGenerosFavoritos cria um novo GenerosFavoritos
func NewGenerosFavoritos(genero, ouvinte string, perfil int) (*GenerosFavoritos, error) {
	gf := &GenerosFavoritos{
		Genero:  genero,
		Perfil:  perfil,
		Ouvinte: ouvinte,
	}
	err := gf.Validate()
	if err != nil {
		return nil, err
	}
	return gf, nil
}

// Validate valida os dados do GenerosFavoritos
func (gf *GenerosFavoritos) Validate() error {
	vld := validator.New()
	if err := vld.Struct(gf); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
