package entity

import (
	"github.com/go-playground/validator/v10"
)

// Album entidade Album
type Album struct {
	Artista       Artista `validate:"required"`
	ID            int     `validate:"required,gte=1"`
	Titulo        string  `validate:"required,gte=1"`
	AnoLancamento int     `validate:"required,gte=1000"`
}

// NewAlbum cria um novo Album
func NewAlbum(email, password, birthday, nomeartistico, biografia, titulo string, anoformacao, anolancamento, id int) (*Album, error) {
	b, err := NewArtista(email, password, birthday, nomeartistico, biografia, anoformacao)
	if err != nil {
		return nil, err
	}
	a := &Album{
		Artista:       *b,
		ID:            id,
		Titulo:        titulo,
		AnoLancamento: anolancamento,
	}
	err = a.Validate()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Validate valida os dados do Album
func (a *Album) Validate() error {
	vld := validator.New()
	if err := vld.Struct(a); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
