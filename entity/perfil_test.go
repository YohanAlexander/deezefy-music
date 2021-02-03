package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerfil(t *testing.T) {

	t.Run("Perfil criado com sucesso", func(t *testing.T) {
		p, err := NewPerfil("syml@spotify.com", "Where is my love", 1)
		assert.Nil(t, err)
		assert.Equal(t, p.InformacoesRelevantes, "Where is my love")
	})

}

func TestPerfil_Validate(t *testing.T) {

	type test struct {
		name                  string
		id                    int
		informacoesrelevantes string
		ouvinte               string
		want                  error
	}

	tests := []test{
		{
			name:                  "Campos válidos",
			id:                    1,
			informacoesrelevantes: "Where is my love",
			ouvinte:               "syml@spotify.com",
			want:                  nil,
		},
		{
			name:                  "Email inválido (user@company.com)",
			id:                    1,
			informacoesrelevantes: "Where is my love",
			ouvinte:               "",
			want:                  ErrInvalidEntity,
		},
		{
			name:                  "InformaçõesRelevantes inválidas",
			id:                    1,
			informacoesrelevantes: "",
			ouvinte:               "syml@spotify.com",
			want:                  ErrInvalidEntity,
		},
		{
			name:                  "ID inválido",
			id:                    0,
			informacoesrelevantes: "Where is my love",
			ouvinte:               "syml@spotify.com",
			want:                  ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPerfil(tc.ouvinte, tc.informacoesrelevantes, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}
