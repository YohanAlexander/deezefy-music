package ouvinte

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewOuvinte(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		o, err := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		assert.Nil(t, err)
		assert.Equal(t, o.PrimeiroNome, "Vance")
	})

}

func TestOuvinte_Validate(t *testing.T) {

	type test struct {
		name         string
		usuario      string
		primeironome string
		sobrenome    string
		want         error
	}

	tests := []test{
		{
			name:         "Campos válidos",
			usuario:      "vancejoy@gmail.com",
			primeironome: "Vance",
			sobrenome:    "Joy",
			want:         nil,
		},
		{
			name:         "PrimeiroNome inválido",
			usuario:      "vancejoy@gmail.com",
			primeironome: "",
			sobrenome:    "Joy",
			want:         entity.ErrInvalidEntity,
		},
		{
			name:         "Sobrenome inválido",
			usuario:      "vancejoy@gmail.com",
			primeironome: "Vance",
			sobrenome:    "",
			want:         entity.ErrInvalidEntity,
		},
		{
			name:         "Email inválido (vance@gmail.com)",
			usuario:      "",
			primeironome: "Vance",
			sobrenome:    "Joy",
			want:         entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewOuvinte(tc.usuario, tc.primeironome, tc.sobrenome)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddTelefone(t *testing.T) {

	t.Run("Telefone criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Telefones))
	})

	t.Run("Telefone já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		err = o.AddTelefone("+5579999999999")
		assert.Equal(t, entity.ErrPhoneRegistered, err)
	})

}

func TestRemoveTelefone(t *testing.T) {

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		err := o.RemoveTelefone("+5579999999999")
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Telefone removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		err := o.RemoveTelefone("+5579999999999")
		assert.Nil(t, err)
	})

}

func TestGetTelefone(t *testing.T) {

	t.Run("Telefone cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		telefone, err := o.GetTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, telefone, "+5579999999999")
	})

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "Vance", "Joy")
		_, err := o.GetTelefone("+5579999999999")
		assert.Equal(t, entity.ErrNotFound, err)
	})

}
