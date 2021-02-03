package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		a, err := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		assert.Nil(t, err)
		assert.Equal(t, a.NomeArtistico, "Vance Joy")
	})

}

func TestArtista_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name          string
		usuario       user
		nomeartistico string
		biografia     string
		anoformacao   int
		want          error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          nil,
		},
		{
			name: "NomeArtistico inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "Biografia inválida",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "AnoFormacao inválido (2000)",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   98,
			want:          ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtista(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.nomeartistico, tc.biografia, tc.anoformacao)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddOuvinte(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Seguidores))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		o, _ = NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err = a.AddOuvinte(*o)
		assert.Equal(t, ErrOuvinteRegistered, err)
	})

}

func TestRemoveOuvinte(t *testing.T) {

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		err := a.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetOuvinte(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		ouvinte, err := a.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_, err := a.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddGravaMusica(t *testing.T) {

	t.Run("Musica criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.AddMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Grava))
	})

	t.Run("Musica já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.AddMusica(*m)
		assert.Nil(t, err)
		m, _ = NewMusica(1, 420, "Creep")
		err = a.AddMusica(*m)
		assert.Equal(t, ErrMusicaRegistered, err)
	})

}

func TestRemoveGravaMusica(t *testing.T) {

	t.Run("Musica não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.RemoveMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Musica removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_ = a.AddMusica(*m)
		err := a.RemoveMusica(*m)
		assert.Nil(t, err)
	})

}

func TestGetGravaMusica(t *testing.T) {

	t.Run("Musica cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_ = a.AddMusica(*m)
		musica, err := a.GetMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, musica, *m)
	})

	t.Run("Musica não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_, err := a.GetMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

}
