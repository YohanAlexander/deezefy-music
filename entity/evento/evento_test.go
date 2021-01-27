package evento

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewEvento(t *testing.T) {

	t.Run("Evento criado com sucesso", func(t *testing.T) {
		e, err := NewEvento("radiohead@gmail.com", "Lollapalooza", 1)
		assert.Nil(t, err)
		assert.Equal(t, e.Nome, "Lollapalooza")
	})

}

func TestEvento_Validate(t *testing.T) {

	type test struct {
		name    string
		usuario string
		nome    string
		id      int
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			usuario: "radiohead@gmail.com",
			nome:    "Lollapalooza",
			id:      1,
			want:    nil,
		},
		{
			name:    "Email inválido (user@company.com)",
			usuario: "",
			nome:    "Lollapalooza",
			id:      1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Nome inválido",
			usuario: "radiohead@gmail.com",
			nome:    "",
			id:      1,
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "ID inválido",
			usuario: "radiohead@gmail.com",
			nome:    "Lollapalooza",
			id:      0,
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewEvento(tc.usuario, tc.nome, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}
