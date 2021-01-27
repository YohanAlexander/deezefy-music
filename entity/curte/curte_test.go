package curte

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewCurte(t *testing.T) {

	t.Run("Curte criada com sucesso", func(t *testing.T) {
		_, err := NewCurte(1, "ouvinte@email.com")
		assert.Nil(t, err)
	})

}

func TestCurte_Validate(t *testing.T) {

	type test struct {
		name    string
		musica  int
		ouvinte string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			musica:  1,
			ouvinte: "ouvinte@email.com",
			want:    nil,
		},
		{
			name:    "Música inválida",
			musica:  0,
			ouvinte: "ouvinte@email.com",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Ouvinte inválido",
			musica:  1,
			ouvinte: "",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCurte(tc.musica, tc.ouvinte)
			assert.Equal(t, err, tc.want)
		})
	}

}
