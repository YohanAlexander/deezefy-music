package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSalvaPlaylist(t *testing.T) {

	t.Run("SalvaPlaylist criada com sucesso", func(t *testing.T) {
		sp, err := NewSalvaPlaylist("Indie Rock", "ouvinte@email.com")
		assert.Nil(t, err)
		assert.Equal(t, sp.Playlist, "Indie Rock")
	})

}

func TestSalvaPlaylist_Validate(t *testing.T) {

	type test struct {
		name     string
		playlist string
		ouvinte  string
		want     error
	}

	tests := []test{
		{
			name:     "Campos válidos",
			playlist: "Indie Rock",
			ouvinte:  "ouvinte@email.com",
			want:     nil,
		},
		{
			name:     "Playlist inválida",
			playlist: "",
			ouvinte:  "ouvinte@email.com",
			want:     ErrInvalidEntity,
		},
		{
			name:     "Ouvinte inválido",
			playlist: "Indie Rock",
			ouvinte:  "",
			want:     ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSalvaPlaylist(tc.playlist, tc.ouvinte)
			assert.Equal(t, err, tc.want)
		})
	}

}
