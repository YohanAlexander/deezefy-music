package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTelefone(t *testing.T) {

	t.Run("Telefone criado com sucesso", func(t *testing.T) {
		_, err := NewTelefone("vancejoy@gmail.com", "+5579999999999")
		assert.Nil(t, err)
	})

}

func TestTelefone_Validate(t *testing.T) {

	type test struct {
		name     string
		ouvinte  string
		telefone string
		want     error
	}

	tests := []test{
		{
			name:     "Campos válidos",
			ouvinte:  "vancejoy@gmail.com",
			telefone: "+5579999999999",
			want:     nil,
		},
		{
			name:     "Telefone inválido (+5579999999999)",
			ouvinte:  "vancejoy@gmail.com",
			telefone: "79999999999",
			want:     ErrInvalidEntity,
		},
		{
			name:     "Email inválido (vance@gmail.com)",
			ouvinte:  "",
			telefone: "+5579999999999",
			want:     ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewTelefone(tc.ouvinte, tc.telefone)
			assert.Equal(t, err, tc.want)
		})
	}

}
