package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMusicaGenero(t *testing.T) {

	t.Run("MusicaGenero criada com sucesso", func(t *testing.T) {
		_, err := NewMusicaGenero("Indie", 1)
		assert.Nil(t, err)
	})

}

func TestMusicaGenero_Validate(t *testing.T) {

	type test struct {
		name   string
		musica int
		genero string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			musica: 1,
			genero: "Indie",
			want:   nil,
		},
		{
			name:   "Musica inválida",
			musica: 0,
			genero: "Indie",
			want:   ErrInvalidEntity,
		},
		{
			name:   "Genero inválido",
			musica: 1,
			genero: "",
			want:   ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMusicaGenero(tc.genero, tc.musica)
			assert.Equal(t, err, tc.want)
		})
	}

}
