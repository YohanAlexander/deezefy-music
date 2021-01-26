package artista

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		a, err := NewArtista("vancejoy@gmail.com", "Vance Joy", "Australian Singer", 2006)
		assert.Nil(t, err)
		assert.Equal(t, a.NomeArtistico, "Vance Joy")
	})

}

func TestArtista_Validate(t *testing.T) {

	type test struct {
		name          string
		usuario       string
		nomeartistico string
		biografia     string
		anoformacao   int
		want          error
	}

	tests := []test{
		{
			name:          "Campos válidos",
			usuario:       "vancejoy@gmail.com",
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          nil,
		},
		{
			name:          "NomeArtistico inválido",
			usuario:       "vancejoy@gmail.com",
			nomeartistico: "",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "Biografia inválida",
			usuario:       "vancejoy@gmail.com",
			nomeartistico: "Vance Joy",
			biografia:     "",
			anoformacao:   2006,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "Email inválido (vance@gmail.com)",
			usuario:       "",
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          entity.ErrInvalidEntity,
		},
		{
			name:          "AnoFormacao inválido (2000)",
			usuario:       "",
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   98,
			want:          entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtista(tc.usuario, tc.nomeartistico, tc.biografia, tc.anoformacao)
			assert.Equal(t, err, tc.want)
		})
	}

}
