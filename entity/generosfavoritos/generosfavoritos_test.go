package generosfavoritos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewGenerosFavoritos(t *testing.T) {

	t.Run("GenerosFavoritos criada com sucesso", func(t *testing.T) {
		_, err := NewGenerosFavoritos("Indie", "ouvinte@email.com", 1)
		assert.Nil(t, err)
	})

}

func TestGenerosFavoritos_Validate(t *testing.T) {

	type test struct {
		name    string
		genero  string
		perfil  int
		ouvinte string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			genero:  "Indie",
			perfil:  1,
			ouvinte: "ouvinte@email.com",
			want:    nil,
		},
		{
			name:    "Ouvinte inválido",
			genero:  "Indie",
			perfil:  1,
			ouvinte: "",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Perfil inválido",
			genero:  "Indie",
			perfil:  0,
			ouvinte: "ouvinte@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Genero inválido",
			genero:  "",
			perfil:  1,
			ouvinte: "ouvinte@email.com",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGenerosFavoritos(tc.genero, tc.ouvinte, tc.perfil)
			assert.Equal(t, err, tc.want)
		})
	}

}
