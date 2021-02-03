package genero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/genero"

	"github.com/stretchr/testify/assert"
)

func newFixtureGenero() *der.Genero {
	return &der.Genero{
		Nome:   "Indie Rock",
		Estilo: "rock",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureGenero()
		_, err := m.CreateGenero(u.Nome, u.Estilo)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureGenero()
	u2 := newFixtureGenero()
	u2.Nome = "Pop Rock"

	email, _ := m.CreateGenero(u1.Nome, u1.Estilo)
	_, _ = m.CreateGenero(u2.Nome, u2.Estilo)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchGeneros("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchGeneros("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListGeneros()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetGenero(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureGenero()
		email, err := m.CreateGenero(u.Nome, u.Estilo)
		assert.Nil(t, err)
		saved, _ := m.GetGenero(email)
		assert.Nil(t, m.UpdateGenero(saved))
		_, err = m.GetGenero(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureGenero()
	u2 := newFixtureGenero()
	u2.Nome = "someone2@deezefy.com"
	email, _ := m.CreateGenero(u2.Nome, u2.Estilo)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteGenero(u1.Nome)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteGenero(email)
		assert.Nil(t, err)
		_, err = m.GetGenero(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureGenero()
		email, _ := m.CreateGenero(u3.Nome, u3.Estilo)
		saved, _ := m.GetGenero(email)
		_ = m.UpdateGenero(saved)
		err = m.DeleteGenero(email)
		assert.Nil(t, err)
	})

}
