package criaplaylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewCriaPlaylist(t *testing.T) {

	t.Run("CriaPlaylist criada com sucesso", func(t *testing.T) {
		_, err := NewCriaPlaylist("2006-01-02", "Indie", "usuario@email.com")
		assert.Nil(t, err)
	})

}

func TestCriaPlaylist_Validate(t *testing.T) {

	type test struct {
		name        string
		datacriacao string
		playlist    string
		usuario     string
		want        error
	}

	tests := []test{
		{
			name:        "Campos válidos",
			datacriacao: "2006-01-02",
			playlist:    "Indie",
			usuario:     "usuario@email.com",
			want:        nil,
		},
		{
			name:        "Usuario inválido",
			datacriacao: "2006-01-02",
			playlist:    "Indie",
			usuario:     "",
			want:        entity.ErrInvalidEntity,
		},
		{
			name:        "Playlist inválida",
			datacriacao: "2006-01-02",
			playlist:    "",
			usuario:     "usuario@email.com",
			want:        entity.ErrInvalidEntity,
		},
		{
			name:        "DataCriação inválida",
			datacriacao: "2006/01/02",
			playlist:    "Indie",
			usuario:     "usuario@email.com",
			want:        entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCriaPlaylist(tc.datacriacao, tc.playlist, tc.usuario)
			assert.Equal(t, err, tc.want)
		})
	}

}
