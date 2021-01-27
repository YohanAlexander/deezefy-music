package musicaplaylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewMusicaPlaylist(t *testing.T) {

	t.Run("MusicaPlaylist criada com sucesso", func(t *testing.T) {
		_, err := NewMusicaPlaylist("Indie Rock", 1)
		assert.Nil(t, err)
	})

}

func TestMusicaPlaylist_Validate(t *testing.T) {

	type test struct {
		name     string
		musica   int
		playlist string
		want     error
	}

	tests := []test{
		{
			name:     "Campos válidos",
			playlist: "Indie Rock",
			musica:   1,
			want:     nil,
		},
		{
			name:     "Musica inválida",
			playlist: "Indie Rock",
			musica:   0,
			want:     entity.ErrInvalidEntity,
		},
		{
			name:     "Playlist inválida",
			playlist: "",
			musica:   1,
			want:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMusicaPlaylist(tc.playlist, tc.musica)
			assert.Equal(t, err, tc.want)
		})
	}

}
