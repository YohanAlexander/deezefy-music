package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEvento(t *testing.T) {

	t.Run("Evento criado com sucesso", func(t *testing.T) {
		e, err := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", "2006-01-02", 1)
		assert.Nil(t, err)
		assert.Equal(t, e.Nome, "Lollapalooza")
	})

}

func TestEvento_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name    string
		usuario user
		id      int
		nome    string
		data    string
		want    error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			id:   1,
			nome: "Lollapalooza",
			data: "2006-01-02",
			want: nil,
		},
		{
			name: "Nome inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			id:   1,
			nome: "",
			data: "2006-01-02",
			want: ErrInvalidEntity,
		},
		{
			name: "ID inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			id:   0,
			nome: "Lollapalooza",
			data: "2006-01-02",
			want: ErrInvalidEntity,
		},
		{
			name: "Data inválida",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			id:   1,
			nome: "Lollapalooza",
			data: "2006/01/02",
			want: ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewEvento(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.nome, tc.data, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}
