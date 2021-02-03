package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsuario(t *testing.T) {

	t.Run("Usuario criado com sucesso", func(t *testing.T) {
		u, err := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		assert.Nil(t, err)
		assert.Equal(t, u.Email, "steve.jobs@apple.com")
		assert.NotEqual(t, u.Password, "new_password")
	})

}

func TestValidatePassword(t *testing.T) {

	t.Run("Senha correta", func(t *testing.T) {
		u, _ := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		err := u.ValidatePassword("new_password")
		assert.Nil(t, err)
	})

	t.Run("Senha incorreta", func(t *testing.T) {
		u, _ := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		err := u.ValidatePassword("wrong_password")
		assert.NotNil(t, err)
	})

}

func TestUsuario_Validate(t *testing.T) {

	type test struct {
		name     string
		email    string
		password string
		birthday string
		want     error
	}

	tests := []test{
		{
			name:     "Campos válidos",
			email:    "steve.jobs@apple.com",
			password: "new_password",
			birthday: "2006-01-02",
			want:     nil,
		},
		{
			name:     "Email inválido (user@company.com)",
			email:    "",
			password: "new_password",
			birthday: "2006-01-02",
			want:     ErrInvalidEntity,
		},
		{
			name:     "Password inválida (12345678)",
			email:    "steve.jobs@apple.com",
			password: "",
			birthday: "2006-01-02",
			want:     ErrInvalidEntity,
		},
		{
			name:     "Birthday inválido (2006-01-02)",
			email:    "steve.jobs@apple.com",
			password: "new_password",
			birthday: "2006/01/02",
			want:     ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewUsuario(tc.email, tc.password, tc.birthday)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddEvento(t *testing.T) {

	t.Run("Evento criado com sucesso", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		err := u.AddEvento(*e)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(u.Organizador))
	})

	t.Run("Evento já registrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		err := u.AddEvento(*e)
		assert.Nil(t, err)
		e, _ = NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		err = u.AddEvento(*e)
		assert.Equal(t, ErrEventoRegistered, err)
	})

}

func TestRemoveEvento(t *testing.T) {

	t.Run("Evento não cadastrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		err := u.RemoveEvento(*e)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Evento removido com sucesso", func(t *testing.T) {
		a, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		_ = a.AddEvento(*e)
		err := a.RemoveEvento(*e)
		assert.Nil(t, err)
	})

}

func TestGetEvento(t *testing.T) {

	t.Run("Evento cadastrado encontrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		_ = u.AddEvento(*e)
		evento, err := u.GetEvento(*e)
		assert.Nil(t, err)
		assert.Equal(t, evento, *e)
	})

	t.Run("Evento não cadastrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		e, _ := NewEvento("radiohead@spotify.com", "somepassword", "2018-02-10", "Lollapalooza", 1)
		_, err := u.GetEvento(*e)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddCriaPlaylist(t *testing.T) {

	t.Run("Playlist criado com sucesso", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		err := u.AddPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(u.Cria))
	})

	t.Run("Playlist já registrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		err := u.AddPlaylist(*p)
		assert.Nil(t, err)
		p, _ = NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		err = u.AddPlaylist(*p)
		assert.Equal(t, ErrPlaylistRegistered, err)
	})

}

func TestRemoveCriaPlaylist(t *testing.T) {

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		err := u.RemovePlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Playlist removido com sucesso", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		_ = u.AddPlaylist(*p)
		err := u.RemovePlaylist(*p)
		assert.Nil(t, err)
	})

}

func TestGetCriaPlaylist(t *testing.T) {

	t.Run("Playlist cadastrado encontrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		_ = u.AddPlaylist(*p)
		playlist, err := u.GetPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, playlist, *p)
	})

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		u, _ := NewUsuario("vancejoy@gmail.com", "somepassword", "2018-02-10")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		_, err := u.GetPlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

}
