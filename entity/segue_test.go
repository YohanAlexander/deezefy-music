package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSegue(t *testing.T) {

	t.Run("Segue criada com sucesso", func(t *testing.T) {
		_, err := NewSegue("artista@email.com", "ouvinte@email.com")
		assert.Nil(t, err)
	})

}

func TestSegue_Validate(t *testing.T) {

	type test struct {
		name    string
		artista string
		ouvinte string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			artista: "artista@email.com",
			ouvinte: "ouvinte@email.com",
			want:    nil,
		},
		{
			name:    "Artista inválido",
			artista: "",
			ouvinte: "ouvinte@email.com",
			want:    ErrInvalidEntity,
		},
		{
			name:    "Ouvinte inválido",
			artista: "artista@email.com",
			ouvinte: "",
			want:    ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSegue(tc.artista, tc.ouvinte)
			assert.Equal(t, err, tc.want)
		})
	}

}
