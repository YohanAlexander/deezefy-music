package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOuvinte(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		o, err := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		assert.Nil(t, err)
		assert.Equal(t, o.PrimeiroNome, "Vance")
	})

}

func TestOuvinte_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name         string
		usuario      user
		primeironome string
		sobrenome    string
		want         error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "Vance",
			sobrenome:    "Joy",
			want:         nil,
		},
		{
			name: "PrimeiroNome inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "",
			sobrenome:    "Joy",
			want:         ErrInvalidEntity,
		},
		{
			name: "Sobrenome inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "Vance",
			sobrenome:    "",
			want:         ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewOuvinte(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.primeironome, tc.sobrenome)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddTelefone(t *testing.T) {

	t.Run("Telefone criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Telefones))
	})

	t.Run("Telefone já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		err = o.AddTelefone("+5579999999999")
		assert.Equal(t, ErrPhoneRegistered, err)
	})

}

func TestRemoveTelefone(t *testing.T) {

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.RemoveTelefone("+5579999999999")
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Telefone removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		err := o.RemoveTelefone("+5579999999999")
		assert.Nil(t, err)
	})

}

func TestGetTelefone(t *testing.T) {

	t.Run("Telefone cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		telefone, err := o.GetTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, telefone, "+5579999999999")
	})

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_, err := o.GetTelefone("+5579999999999")
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.AddArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Seguindo))
	})

	t.Run("Artista já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.AddArtista(*a)
		assert.Nil(t, err)
		a, _ = NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err = o.AddArtista(*a)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveArtista(t *testing.T) {

	t.Run("Artista não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.RemoveArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Artista removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = o.AddArtista(*a)
		err := o.RemoveArtista(*a)
		assert.Nil(t, err)
	})

}

func TestGetArtista(t *testing.T) {

	t.Run("Artista cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = o.AddArtista(*a)
		artista, err := o.GetArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, artista, *a)
	})

	t.Run("Artista não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_, err := o.GetArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}
