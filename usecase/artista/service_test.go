package artista

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artista"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/artista"

	"github.com/stretchr/testify/assert"
)

func newFixtureArtista() *der.Artista {
	return &der.Artista{
		Usuario:       "someone@deezefy.com",
		NomeArtistico: "Imagine Dragons",
		Biografia:     "Indie Rock Band",
		AnoFormacao:   1998,
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtista()
		_, err := m.CreateArtista(u.Usuario, u.NomeArtistico, u.Biografia, u.AnoFormacao)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtista()
	u2 := newFixtureArtista()
	u2.Usuario = "someone2@deezefy.com"

	email, _ := m.CreateArtista(u1.Usuario, u1.NomeArtistico, u1.Biografia, u1.AnoFormacao)
	_, _ = m.CreateArtista(u2.Usuario, u2.NomeArtistico, u2.Biografia, u2.AnoFormacao)

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
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtista()
		email, err := m.CreateArtista(u.Usuario, u.NomeArtistico, u.Biografia, u.AnoFormacao)
		assert.Nil(t, err)
		saved, _ := m.GetArtista(email)
		assert.Nil(t, m.UpdateArtista(saved))
		_, err = m.GetArtista(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtista()
	u2 := newFixtureArtista()
	u2.Usuario = "someone2@deezefy.com"
	email, _ := m.CreateArtista(u2.Usuario, u2.NomeArtistico, u2.Biografia, u2.AnoFormacao)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteArtista(u1.Usuario)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteArtista(email)
		assert.Nil(t, err)
		_, err = m.GetArtista(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureArtista()
		email, _ := m.CreateArtista(u3.Usuario, u3.NomeArtistico, u3.Biografia, u3.AnoFormacao)
		saved, _ := m.GetArtista(email)
		_ = m.UpdateArtista(saved)
		err = m.DeleteArtista(email)
		assert.Nil(t, err)
	})

}
