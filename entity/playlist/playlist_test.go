package playlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewPlaylist(t *testing.T) {

	t.Run("Playlist criada com sucesso", func(t *testing.T) {
		p, err := NewPlaylist("Indie Rock", "ativo")
		assert.Nil(t, err)
		assert.Equal(t, p.Status, "ativo")
	})

}

func TestPlaylist_Validate(t *testing.T) {

	type test struct {
		name   string
		nome   string
		status string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			nome:   "Indie Rock",
			status: "ativo",
			want:   nil,
		},
		{
			name:   "Nome inválido",
			nome:   "",
			status: "ativo",
			want:   entity.ErrInvalidEntity,
		},
		{
			name:   "Status inválido",
			nome:   "Indie Rock",
			status: "ativa",
			want:   entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPlaylist(tc.nome, tc.status)
			assert.Equal(t, err, tc.want)
		})
	}

}
