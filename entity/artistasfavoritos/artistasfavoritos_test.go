package artistasfavoritos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewArtistasFavoritos(t *testing.T) {

	t.Run("ArtistasFavoritos criada com sucesso", func(t *testing.T) {
		_, err := NewArtistasFavoritos(1, "ouvinte@email.com", "artista@email.com")
		assert.Nil(t, err)
	})

}

func TestArtistasFavoritos_Validate(t *testing.T) {

	type test struct {
		name    string
		perfil  int
		ouvinte string
		artista string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			perfil:  1,
			ouvinte: "ouvinte@email.com",
			artista: "artista@email.com",
			want:    nil,
		},
		{
			name:    "Perfil inválido",
			perfil:  0,
			ouvinte: "ouvinte@email.com",
			artista: "artista@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Ouvinte inválido",
			perfil:  1,
			ouvinte: "",
			artista: "artista@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			perfil:  1,
			ouvinte: "ouvinte@email.com",
			artista: "",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtistasFavoritos(tc.perfil, tc.ouvinte, tc.artista)
			assert.Equal(t, err, tc.want)
		})
	}

}
