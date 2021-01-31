package ouvinte

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ouvinte"

	"github.com/stretchr/testify/assert"
)

func newFixtureOuvinte() *der.Ouvinte {
	return &der.Ouvinte{
		Usuario:      "someone@deezefy.com",
		PrimeiroNome: "Imagine",
		Sobrenome:    "Dragons",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureOuvinte()
		_, err := m.CreateOuvinte(u.Usuario, u.PrimeiroNome, u.Sobrenome)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureOuvinte()
	u2 := newFixtureOuvinte()
	u2.Usuario = "someone2@deezefy.com"

	email, _ := m.CreateOuvinte(u1.Usuario, u1.PrimeiroNome, u1.Sobrenome)
	_, _ = m.CreateOuvinte(u2.Usuario, u2.PrimeiroNome, u2.Sobrenome)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchOuvintes("someone@deezefy.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchOuvintes("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListOuvintes()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetOuvinte(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureOuvinte()
		email, err := m.CreateOuvinte(u.Usuario, u.PrimeiroNome, u.Sobrenome)
		assert.Nil(t, err)
		saved, _ := m.GetOuvinte(email)
		assert.Nil(t, m.UpdateOuvinte(saved))
		_, err = m.GetOuvinte(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureOuvinte()
	u2 := newFixtureOuvinte()
	u2.Usuario = "someone2@deezefy.com"
	email, _ := m.CreateOuvinte(u2.Usuario, u2.PrimeiroNome, u2.Sobrenome)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteOuvinte(u1.Usuario)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteOuvinte(email)
		assert.Nil(t, err)
		_, err = m.GetOuvinte(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureOuvinte()
		email, _ := m.CreateOuvinte(u3.Usuario, u3.PrimeiroNome, u3.Sobrenome)
		saved, _ := m.GetOuvinte(email)
		_ = m.UpdateOuvinte(saved)
		err = m.DeleteOuvinte(email)
		assert.Nil(t, err)
	})

}
