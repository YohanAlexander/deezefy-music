package grava

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewGrava(t *testing.T) {

	t.Run("Grava criada com sucesso", func(t *testing.T) {
		_, err := NewGrava(1, "artista@email.com")
		assert.Nil(t, err)
	})

}

func TestGrava_Validate(t *testing.T) {

	type test struct {
		name    string
		musica  int
		artista string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			musica:  1,
			artista: "artista@email.com",
			want:    nil,
		},
		{
			name:    "Música inválida",
			musica:  0,
			artista: "artista@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			musica:  1,
			artista: "",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGrava(tc.musica, tc.artista)
			assert.Equal(t, err, tc.want)
		})
	}

}
