package entity

import (
	"github.com/go-playground/validator/v10"
)

// SalvaPlaylist entidade SalvaPlaylist
type SalvaPlaylist struct {
	Playlist string `validate:"required,gte=1"`
	Ouvinte  string `validate:"required,email"`
}

// NewSalvaPlaylist cria um novo SalvaPlaylist
func NewSalvaPlaylist(playlist, ouvinte string) (*SalvaPlaylist, error) {
	sp := &SalvaPlaylist{
		Playlist: playlist,
		Ouvinte:  ouvinte,
	}
	err := sp.Validate()
	if err != nil {
		return nil, err
	}
	return sp, nil
}

// Validate valida os dados do SalvaPlaylist
func (sp *SalvaPlaylist) Validate() error {
	vld := validator.New()
	if err := vld.Struct(sp); err != nil {
		return ErrInvalidEntity
	}
	return nil
}
