package local

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewLocal(t *testing.T) {

	t.Run("Local criado com sucesso", func(t *testing.T) {
		l, err := NewLocal("São Paulo", "Brazil", 1)
		assert.Nil(t, err)
		assert.Equal(t, l.Pais, "Brazil")
	})

}

func TestLocal_Validate(t *testing.T) {

	type test struct {
		name   string
		id     int
		cidade string
		pais   string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			id:     1,
			cidade: "São Paulo",
			pais:   "Brazil",
			want:   nil,
		},
		{
			name:   "Cidade inválida",
			id:     1,
			cidade: "",
			pais:   "Brazil",
			want:   entity.ErrInvalidEntity,
		},
		{
			name:   "Pais inválido",
			id:     1,
			cidade: "São Paulo",
			pais:   "",
			want:   entity.ErrInvalidEntity,
		},
		{
			name:   "ID inválido",
			id:     0,
			cidade: "São Paulo",
			pais:   "Brazil",
			want:   entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewLocal(tc.cidade, tc.pais, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}
