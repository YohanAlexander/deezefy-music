package salvaalbum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewSalvaAlbum(t *testing.T) {

	t.Run("SalvaAlbum criada com sucesso", func(t *testing.T) {
		_, err := NewSalvaAlbum(1, "ouvinte@email.com", "artista@email.com")
		assert.Nil(t, err)
	})

}

func TestSalvaAlbum_Validate(t *testing.T) {

	type test struct {
		name    string
		album   int
		ouvinte string
		artista string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			album:   1,
			ouvinte: "ouvinte@email.com",
			artista: "artista@email.com",
			want:    nil,
		},
		{
			name:    "Album inválido",
			album:   0,
			ouvinte: "ouvinte@email.com",
			artista: "artista@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Ouvinte inválido",
			album:   1,
			ouvinte: "",
			artista: "artista@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			album:   1,
			ouvinte: "ouvinte@email.com",
			artista: "",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSalvaAlbum(tc.album, tc.ouvinte, tc.artista)
			assert.Equal(t, err, tc.want)
		})
	}

}
