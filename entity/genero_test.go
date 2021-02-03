package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenero(t *testing.T) {

	t.Run("Genero criado com sucesso", func(t *testing.T) {
		g, err := NewGenero("Indie Rock", "rock")
		assert.Nil(t, err)
		assert.Equal(t, g.Estilo, "rock")
	})

}

func TestGenero_Validate(t *testing.T) {

	type test struct {
		name   string
		nome   string
		estilo string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			nome:   "Indie Rock",
			estilo: "rock",
			want:   nil,
		},
		{
			name:   "Nome inválido",
			nome:   "",
			estilo: "rock",
			want:   ErrInvalidEntity,
		},
		{
			name:   "Estilo inválido",
			nome:   "Indie Rock",
			estilo: "indie",
			want:   ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGenero(tc.nome, tc.estilo)
			assert.Equal(t, err, tc.want)
		})
	}

}
