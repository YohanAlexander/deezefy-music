package artista

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureArtista() *entity.Artista {
	return &entity.Artista{
		Usuario: entity.Usuario{
			Email:    "someone@deezefy.com",
			Password: "12345678",
			Birthday: "1998-05-27",
		},
		NomeArtistico: "Imagine Dragons",
		Biografia:     "Indie Rock Band",
		AnoFormacao:   1998,
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		a := newFixtureArtista()
		_, err := m.CreateArtista(a.Usuario.Email, a.Usuario.Password, a.Usuario.Birthday, a.NomeArtistico, a.Biografia, a.AnoFormacao)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	a1 := newFixtureArtista()
	a2 := newFixtureArtista()
	a2.Usuario = entity.Usuario{
		Email:    "someone2@deezefy.com",
		Password: "12345678",
		Birthday: "1998-05-27",
	}

	email, _ := m.CreateArtista(a1.Usuario.Email, a1.Usuario.Password, a1.Usuario.Birthday, a1.NomeArtistico, a1.Biografia, a1.AnoFormacao)
	_, _ = m.CreateArtista(a2.Usuario.Email, a2.Usuario.Password, a2.Usuario.Birthday, a2.NomeArtistico, a2.Biografia, a2.AnoFormacao)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchArtistas("someone@deezefy.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchArtistas("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListArtistas()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetArtista(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		a := newFixtureArtista()
		email, err := m.CreateArtista(a.Usuario.Email, a.Usuario.Password, a.Usuario.Birthday, a.NomeArtistico, a.Biografia, a.AnoFormacao)
		assert.Nil(t, err)
		saved, _ := m.GetArtista(email)
		assert.Nil(t, m.UpdateArtista(saved))
		_, err = m.GetArtista(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	a1 := newFixtureArtista()
	a2 := newFixtureArtista()
	a2.Usuario = entity.Usuario{
		Email:    "someone2@deezefy.com",
		Password: "12345678",
		Birthday: "1998-05-27",
	}
	email, _ := m.CreateArtista(a2.Usuario.Email, a2.Usuario.Password, a2.Usuario.Birthday, a2.NomeArtistico, a2.Biografia, a2.AnoFormacao)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteArtista(a1.Usuario.Email)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteArtista(email)
		assert.Nil(t, err)
		_, err = m.GetArtista(email)
		assert.Equal(t, entity.ErrNotFound, err)

		a3 := newFixtureArtista()
		email, _ := m.CreateArtista(a3.Usuario.Email, a3.Usuario.Password, a3.Usuario.Birthday, a3.NomeArtistico, a3.Biografia, a3.AnoFormacao)
		saved, _ := m.GetArtista(email)
		_ = m.UpdateArtista(saved)
		err = m.DeleteArtista(email)
		assert.Nil(t, err)
	})

}
