package segue

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/segue"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/segue"

	"github.com/stretchr/testify/assert"
)

func newFixtureSegue() *der.Segue {
	return &der.Segue{
		Artista: "artista@email.com",
		Ouvinte: "ouvinte@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSegue()
		_, _, err := m.CreateSegue(u.Artista, u.Ouvinte)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSegue()
	u2 := newFixtureSegue()
	u2.Artista = "artista2@email.com"
	u2.Ouvinte = "ouvinte2@email.com"

	artista, ouvinte, _ := m.CreateSegue(u1.Artista, u1.Ouvinte)
	_, _, _ = m.CreateSegue(u2.Artista, u2.Ouvinte)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchSegues("artista@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchSegues("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListSegues()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetSegue(artista, ouvinte)
		assert.Nil(t, err)
	})

	t.Run("get by artista", func(t *testing.T) {
		_, err := m.GetSegueByArtista(artista)
		assert.NotNil(t, err)
	})

	t.Run("get by ouvinte", func(t *testing.T) {
		_, err := m.GetSegueByOuvinte(ouvinte)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSegue()
		artista, ouvinte, err := m.CreateSegue(u.Artista, u.Ouvinte)
		assert.Nil(t, err)
		saved, _ := m.GetSegue(artista, ouvinte)
		assert.Nil(t, m.UpdateSegue(saved))
		_, err = m.GetSegue(artista, ouvinte)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSegue()
	u2 := newFixtureSegue()
	u2.Artista = "artista2@email.com"
	u2.Ouvinte = "ouvinte2@email.com"
	artista, ouvinte, _ := m.CreateSegue(u2.Artista, u2.Ouvinte)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteSegue(u1.Artista, u1.Ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteSegue(artista, ouvinte)
		assert.Nil(t, err)
		_, err = m.GetSegue(artista, ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureSegue()
		artista, ouvinte, _ := m.CreateSegue(u3.Artista, u3.Ouvinte)
		saved, _ := m.GetSegue(artista, ouvinte)
		_ = m.UpdateSegue(saved)
		err = m.DeleteSegue(artista, ouvinte)
		assert.Nil(t, err)
	})

}
