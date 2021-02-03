package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerfil(t *testing.T) {

	t.Run("Perfil criado com sucesso", func(t *testing.T) {
		p, err := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		assert.Nil(t, err)
		assert.Equal(t, p.InformacoesRelevantes, "Where is my love")
	})

}

func TestPerfil_Validate(t *testing.T) {

	type ouvinte struct {
		email        string
		password     string
		birthday     string
		primeironome string
		sobrenome    string
	}

	type test struct {
		name                  string
		ouvinte               ouvinte
		id                    int
		informacoesrelevantes string
		want                  error
	}

	tests := []test{
		{
			name:                  "Campos válidos",
			informacoesrelevantes: "Where is my love",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   1,
			want: nil,
		},
		{
			name:                  "InformaçõesRelevantes inválidas",
			informacoesrelevantes: "",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   1,
			want: ErrInvalidEntity,
		},
		{
			name:                  "ID inválido",
			informacoesrelevantes: "Where is my love",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   0,
			want: ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPerfil(tc.ouvinte.email, tc.ouvinte.password, tc.ouvinte.birthday, tc.ouvinte.primeironome, tc.ouvinte.sobrenome, tc.informacoesrelevantes, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}
