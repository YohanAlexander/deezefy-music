package albummusica

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewAlbumMusica(t *testing.T) {

	t.Run("AlbumMusica criada com sucesso", func(t *testing.T) {
		_, err := NewAlbumMusica("artista@email.com", 1, 1)
		assert.Nil(t, err)
	})

}

func TestAlbumMusica_Validate(t *testing.T) {

	type test struct {
		name    string
		artista string
		album   int
		musica  int
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			artista: "artista@email.com",
			album:   1,
			musica:  1,
			want:    nil,
		},
		{
			name:    "Musica inválida",
			artista: "artista@email.com",
			album:   1,
			musica:  0,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Album inválido",
			artista: "artista@email.com",
			album:   0,
			musica:  1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			artista: "",
			album:   1,
			musica:  1,
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAlbumMusica(tc.artista, tc.album, tc.musica)
			assert.Equal(t, err, tc.want)
		})
	}

}
