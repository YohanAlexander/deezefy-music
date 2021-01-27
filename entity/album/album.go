package album

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// Album entidade Album
type Album struct {
	Artista       string `validate:"required,email"`
	ID            int    `validate:"required,gte=1"`
	Titulo        string `validate:"required,gte=1"`
	AnoLancamento int    `validate:"required,gte=1900"`
}

// NewAlbum cria um novo Album
func NewAlbum(id, anolancamento int, titulo, artista string) (*Album, error) {
	a := &Album{
		ID:            id,
		AnoLancamento: anolancamento,
		Titulo:        titulo,
		Artista:       artista,
	}
	err := a.Validate()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Validate valida os dados do Album
func (a *Album) Validate() error {
	vld := validator.New()
	if err := vld.Struct(a); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
