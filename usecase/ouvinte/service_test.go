package ouvinte

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureOuvinte() *entity.Ouvinte {
	return &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email:    "someone@deezefy.com",
			Password: "12345678",
			Birthday: "1998-05-27",
		},
		PrimeiroNome: "Imagine",
		Sobrenome:    "Dragons",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		o := newFixtureOuvinte()
		_, err := m.CreateOuvinte(o.Usuario.Email, o.Usuario.Password, o.Usuario.Birthday, o.PrimeiroNome, o.Sobrenome)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	o1 := newFixtureOuvinte()
	o2 := newFixtureOuvinte()
	o2.Usuario = entity.Usuario{
		Email:    "someone2@deezefy.com",
		Password: "12345678",
		Birthday: "1998-05-27",
	}

	email, _ := m.CreateOuvinte(o1.Usuario.Email, o1.Usuario.Password, o1.Usuario.Birthday, o1.PrimeiroNome, o1.Sobrenome)
	_, _ = m.CreateOuvinte(o2.Usuario.Email, o2.Usuario.Password, o2.Usuario.Birthday, o2.PrimeiroNome, o2.Sobrenome)

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
		o := newFixtureOuvinte()
		email, err := m.CreateOuvinte(o.Usuario.Email, o.Usuario.Password, o.Usuario.Birthday, o.PrimeiroNome, o.Sobrenome)
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
	o1 := newFixtureOuvinte()
	o2 := newFixtureOuvinte()
	o2.Usuario = entity.Usuario{
		Email:    "someone2@deezefy.com",
		Password: "12345678",
		Birthday: "1998-05-27",
	}
	email, _ := m.CreateOuvinte(o2.Usuario.Email, o2.Usuario.Password, o2.Usuario.Birthday, o2.PrimeiroNome, o2.Sobrenome)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteOuvinte(o1.Usuario.Email)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteOuvinte(email)
		assert.Nil(t, err)
		_, err = m.GetOuvinte(email)
		assert.Equal(t, entity.ErrNotFound, err)

		o3 := newFixtureOuvinte()
		email, _ := m.CreateOuvinte(o3.Usuario.Email, o3.Usuario.Password, o3.Usuario.Birthday, o3.PrimeiroNome, o3.Sobrenome)
		saved, _ := m.GetOuvinte(email)
		_ = m.UpdateOuvinte(saved)
		err = m.DeleteOuvinte(email)
		assert.Nil(t, err)
	})

}
