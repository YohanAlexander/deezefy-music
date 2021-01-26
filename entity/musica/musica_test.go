package musica

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewMusica(t *testing.T) {

	t.Run("Musica criada com sucesso", func(t *testing.T) {
		m, err := NewMusica(1, 420, "Creep")
		assert.Nil(t, err)
		assert.Equal(t, m.Nome, "Creep")
	})

}

func TestMusica_Validate(t *testing.T) {

	type test struct {
		name    string
		id      int
		duracao int
		nome    string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			id:      1,
			duracao: 100,
			nome:    "Creep",
			want:    nil,
		},
		{
			name:    "ID inválido",
			id:      0,
			duracao: 100,
			nome:    "Creep",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Nome inválido",
			id:      1,
			duracao: 100,
			nome:    "",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "Duração inválida",
			id:      1,
			duracao: 10,
			nome:    "Creep",
			want:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMusica(tc.id, tc.duracao, tc.nome)
			assert.Equal(t, err, tc.want)
		})
	}

}
