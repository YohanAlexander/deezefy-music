package entity

import (
	"github.com/go-playground/validator/v10"
)

// Playlist entidade Playlist
type Playlist struct {
	Nome   string `validate:"required,gte=1"`
	Status string `validate:"required,oneof=ativo inativo"`
}

// NewPlaylist cria um novo Playlist
func NewPlaylist(nome, status string) (*Playlist, error) {
	p := &Playlist{
		Nome:   nome,
		Status: status,
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate valida os dados do Playlist
func (p *Playlist) Validate() error {
	vld := validator.New()
	if err := vld.Struct(p); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
