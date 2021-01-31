package evento

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/evento"

	"github.com/stretchr/testify/assert"
)

func newFixtureEvento() *der.Evento {
	return &der.Evento{
		ID:      815,
		Nome:    "Lollapalooza",
		Usuario: "someone@spotify.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureEvento()
		_, err := m.CreateEvento(u.Usuario, u.Nome, u.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureEvento()
	u2 := newFixtureEvento()
	u2.ID = 200
	u2.Nome = "Rock in Rio"

	email, _ := m.CreateEvento(u1.Usuario, u1.Nome, u1.ID)
	_, _ = m.CreateEvento(u2.Usuario, u2.Nome, u2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchEventos("Lollapalooza")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchEventos("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListEventos()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetEvento(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureEvento()
		email, err := m.CreateEvento(u.Usuario, u.Nome, u.ID)
		assert.Nil(t, err)
		saved, _ := m.GetEvento(email)
		assert.Nil(t, m.UpdateEvento(saved))
		_, err = m.GetEvento(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureEvento()
	u2 := newFixtureEvento()
	u2.ID = 200
	email, _ := m.CreateEvento(u2.Usuario, u2.Nome, u2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteEvento(u1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteEvento(email)
		assert.Nil(t, err)
		_, err = m.GetEvento(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureEvento()
		email, _ := m.CreateEvento(u3.Usuario, u3.Nome, u3.ID)
		saved, _ := m.GetEvento(email)
		_ = m.UpdateEvento(saved)
		err = m.DeleteEvento(email)
		assert.Nil(t, err)
	})

}
