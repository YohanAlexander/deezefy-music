package album

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewAlbum(t *testing.T) {

	t.Run("Album criado com sucesso", func(t *testing.T) {
		a, err := NewAlbum(1, 2000, "Viva la vida", "coldplay@gmail.com")
		assert.Nil(t, err)
		assert.Equal(t, a.Titulo, "Viva la vida")
	})

}

func TestAlbum_Validate(t *testing.T) {

	type test struct {
		name          string
		artista       string
		id            int
		titulo        string
		anolancamento int
		want          error
	}

	tests := []test{
		{
			name:          "Campos válidos",
			artista:       "coldplay@gmail.com",
			id:            1,
			titulo:        "Yellow",
			anolancamento: 2000,
			want:          nil,
		},
		{
			name:          "Email inválido (user@company.com)",
			artista:       "",
			id:            1,
			titulo:        "Yellow",
			anolancamento: 2000,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "ID inválido",
			artista:       "coldplay@gmail.com",
			id:            0,
			titulo:        "Yellow",
			anolancamento: 2000,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "Título inválido",
			artista:       "coldplay@gmail.com",
			id:            1,
			titulo:        "",
			anolancamento: 2000,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "AnoLancamento inválido",
			artista:       "coldplay@gmail.com",
			id:            1,
			titulo:        "Yellow",
			anolancamento: 200,
			want:          entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAlbum(tc.id, tc.anolancamento, tc.titulo, tc.artista)
			assert.Equal(t, err, tc.want)
		})
	}

}
