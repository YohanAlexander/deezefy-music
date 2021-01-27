package ocorre

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewOcorre(t *testing.T) {

	t.Run("Ocorre criada com sucesso", func(t *testing.T) {
		_, err := NewOcorre("2006-01-02", "artista@email.com", "usuario@email.com", 1, 1)
		assert.Nil(t, err)
	})

}

func TestOcorre_Validate(t *testing.T) {

	type test struct {
		name    string
		data    string
		artista string
		usuario string
		local   int
		evento  int
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			data:    "2006-01-02",
			artista: "artista@email.com",
			usuario: "usuario@email.com",
			local:   1,
			evento:  1,
			want:    nil,
		},
		{
			name:    "Data inválida",
			data:    "2006/01/02",
			artista: "artista@email.com",
			usuario: "usuario@email.com",
			local:   1,
			evento:  1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Artista inválido",
			data:    "2006-01-02",
			artista: "",
			usuario: "usuario@email.com",
			local:   1,
			evento:  1,
			want:    entity.ErrInvalidEntity,
		}, {
			name:    "Usuario inválido",
			data:    "2006-01-02",
			artista: "artista@email.com",
			usuario: "",
			local:   1,
			evento:  1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Local inválido",
			data:    "2006-01-02",
			artista: "artista@email.com",
			usuario: "usuario@email.com",
			local:   0,
			evento:  1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Evento inválido",
			data:    "2006-01-02",
			artista: "artista@email.com",
			usuario: "usuario@email.com",
			local:   1,
			evento:  0,
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewOcorre(tc.data, tc.artista, tc.usuario, tc.local, tc.evento)
			assert.Equal(t, err, tc.want)
		})
	}

}
