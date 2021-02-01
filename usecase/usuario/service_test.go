package usuario

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/usuario"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/usuario"

	"github.com/stretchr/testify/assert"
)

func newFixtureUsuario() *der.Usuario {
	return &der.Usuario{
		Email:    "someone@deezefy.com",
		Password: "12345678",
		Birthday: "1998-05-27",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureUsuario()
		_, err := m.CreateUsuario(u.Email, u.Password, u.Birthday)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureUsuario()
	u2 := newFixtureUsuario()
	u2.Email = "someone2@deezefy.com"

	email, _ := m.CreateUsuario(u1.Email, u1.Password, u1.Birthday)
	_, _ = m.CreateUsuario(u2.Email, u2.Password, u2.Birthday)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchUsuarios("someone@deezefy.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchUsuarios("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListUsuarios()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetUsuario(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureUsuario()
		email, err := m.CreateUsuario(u.Email, u.Password, u.Birthday)
		assert.Nil(t, err)
		saved, _ := m.GetUsuario(email)
		assert.Nil(t, m.UpdateUsuario(saved))
		_, err = m.GetUsuario(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureUsuario()
	u2 := newFixtureUsuario()
	u2.Email = "someone2@deezefy.com"
	email, _ := m.CreateUsuario(u2.Email, u2.Password, u2.Birthday)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteUsuario(u1.Email)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteUsuario(email)
		assert.Nil(t, err)
		_, err = m.GetUsuario(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureUsuario()
		email, _ := m.CreateUsuario(u3.Email, u3.Password, u3.Birthday)
		saved, _ := m.GetUsuario(email)
		_ = m.UpdateUsuario(saved)
		err = m.DeleteUsuario(email)
		assert.Nil(t, err)
	})

}
