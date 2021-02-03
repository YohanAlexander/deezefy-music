package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		a, err := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		assert.Nil(t, err)
		assert.Equal(t, a.NomeArtistico, "Vance Joy")
	})

}

func TestArtista_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name          string
		usuario       user
		nomeartistico string
		biografia     string
		anoformacao   int
		want          error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          nil,
		},
		{
			name: "NomeArtistico inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "Biografia inválida",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "AnoFormacao inválido (2000)",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   98,
			want:          ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtista(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.nomeartistico, tc.biografia, tc.anoformacao)
			assert.Equal(t, err, tc.want)
		})
	}

}
