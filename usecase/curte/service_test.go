package curte

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/curte"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/curte"

	"github.com/stretchr/testify/assert"
)

func newFixtureCurte() *der.Curte {
	return &der.Curte{
		Musica:  1,
		Ouvinte: "ouvinte@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureCurte()
		_, _, err := m.CreateCurte(u.Musica, u.Ouvinte)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureCurte()
	u2 := newFixtureCurte()
	u2.Musica = 2
	u2.Ouvinte = "ouvinte2@email.com"

	musica, ouvinte, _ := m.CreateCurte(u1.Musica, u1.Ouvinte)
	_, _, _ = m.CreateCurte(u2.Musica, u2.Ouvinte)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchCurtes("ouvinte@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchCurtes("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListCurtes()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetCurte(musica, ouvinte)
		assert.Nil(t, err)
	})

	t.Run("get by Musica", func(t *testing.T) {
		_, err := m.GetCurteByMusica(musica)
		assert.NotNil(t, err)
	})

	t.Run("get by ouvinte", func(t *testing.T) {
		_, err := m.GetCurteByOuvinte(ouvinte)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureCurte()
		musica, ouvinte, err := m.CreateCurte(u.Musica, u.Ouvinte)
		assert.Nil(t, err)
		saved, _ := m.GetCurte(musica, ouvinte)
		assert.Nil(t, m.UpdateCurte(saved))
		_, err = m.GetCurte(musica, ouvinte)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureCurte()
	u2 := newFixtureCurte()
	u2.Musica = 2
	u2.Ouvinte = "ouvinte2@email.com"
	musica, ouvinte, _ := m.CreateCurte(u2.Musica, u2.Ouvinte)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteCurte(u1.Musica, u1.Ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteCurte(musica, ouvinte)
		assert.Nil(t, err)
		_, err = m.GetCurte(musica, ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureCurte()
		musica, ouvinte, _ := m.CreateCurte(u3.Musica, u3.Ouvinte)
		saved, _ := m.GetCurte(musica, ouvinte)
		_ = m.UpdateCurte(saved)
		err = m.DeleteCurte(musica, ouvinte)
		assert.Nil(t, err)
	})

}
