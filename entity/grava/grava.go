package grava

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// Grava entidade Grava
type Grava struct {
	Musica  int    `validate:"required,gte=1"`
	Artista string `validate:"required,email"`
}

// NewGrava cria um novo Grava
func NewGrava(musica int, artista string) (*Grava, error) {
	g := &Grava{
		Musica:  musica,
		Artista: artista,
	}
	err := g.Validate()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Validate valida os dados do Grava
func (g *Grava) Validate() error {
	vld := validator.New()
	if err := vld.Struct(g); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
