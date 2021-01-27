package criaplaylist

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// CriaPlaylist entidade CriaPlaylist
type CriaPlaylist struct {
	DataCriacao string `validate:"required,datetime=2006-01-02"`
	Playlist    string `validate:"required,gte=1"`
	Usuario     string `validate:"required,email"`
}

// NewCriaPlaylist cria um novo CriaPlaylist
func NewCriaPlaylist(datacriacao, playlist, usuario string) (*CriaPlaylist, error) {
	cp := &CriaPlaylist{
		DataCriacao: datacriacao,
		Playlist:    playlist,
		Usuario:     usuario,
	}
	err := cp.Validate()
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Validate valida os dados do CriaPlaylist
func (cp *CriaPlaylist) Validate() error {
	vld := validator.New()
	if err := vld.Struct(cp); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
