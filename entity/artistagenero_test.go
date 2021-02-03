package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtistaGenero(t *testing.T) {

	t.Run("ArtistaGenero criada com sucesso", func(t *testing.T) {
		_, err := NewArtistaGenero("artista@email.com", "Indie")
		assert.Nil(t, err)
	})

}

func TestArtistaGenero_Validate(t *testing.T) {

	type test struct {
		name    string
		artista string
		genero  string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			artista: "artista@email.com",
			genero:  "Funk",
			want:    nil,
		},
		{
			name:    "Genero inválido",
			artista: "artista@email.com",
			genero:  "",
			want:    ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			artista: "",
			genero:  "Funk",
			want:    ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtistaGenero(tc.artista, tc.genero)
			assert.Equal(t, err, tc.want)
		})
	}

}
