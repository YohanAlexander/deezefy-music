package musicaplaylist

import (
	"github.com/go-playground/validator/v10"
	"github.com/yohanalexander/deezefy-music/entity"
)

// MusicaPlaylist entidade MusicaPlaylist
type MusicaPlaylist struct {
	Musica   int    `validate:"required,gte=1"`
	Playlist string `validate:"required,gte=1"`
}

// NewMusicaPlaylist cria um novo MusicaPlaylist
func NewMusicaPlaylist(playlist string, musica int) (*MusicaPlaylist, error) {
	mp := &MusicaPlaylist{
		Musica:   musica,
		Playlist: playlist,
	}
	err := mp.Validate()
	if err != nil {
		return nil, err
	}
	return mp, nil
}

// Validate valida os dados do MusicaPlaylist
func (mp *MusicaPlaylist) Validate() error {
	vld := validator.New()
	if err := vld.Struct(mp); err != nil {
		return entity.ErrInvalidEntity
	}
	return nil
}
