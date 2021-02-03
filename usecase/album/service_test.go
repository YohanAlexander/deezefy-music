package album

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureAlbum() *entity.Album {
	return &entity.Album{
		ID:            815,
		AnoLancamento: 1998,
		Titulo:        "Cage The Elephant",
		Artista:       "someone@spotify.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureAlbum()
		_, err := m.CreateAlbum(u.ID, u.AnoLancamento, u.Titulo, u.Artista)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureAlbum()
	u2 := newFixtureAlbum()
	u2.ID = 200
	u2.Titulo = "Radioactive"

	email, _ := m.CreateAlbum(u1.ID, u1.AnoLancamento, u1.Titulo, u1.Artista)
	_, _ = m.CreateAlbum(u2.ID, u2.AnoLancamento, u2.Titulo, u2.Artista)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchAlbums("Cage The Elephant")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchAlbums("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListAlbums()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetAlbum(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureAlbum()
		email, err := m.CreateAlbum(u.ID, u.AnoLancamento, u.Titulo, u.Artista)
		assert.Nil(t, err)
		saved, _ := m.GetAlbum(email)
		assert.Nil(t, m.UpdateAlbum(saved))
		_, err = m.GetAlbum(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureAlbum()
	u2 := newFixtureAlbum()
	u2.ID = 200
	email, _ := m.CreateAlbum(u2.ID, u2.AnoLancamento, u2.Titulo, u2.Artista)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteAlbum(u1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteAlbum(email)
		assert.Nil(t, err)
		_, err = m.GetAlbum(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureAlbum()
		email, _ := m.CreateAlbum(u3.ID, u3.AnoLancamento, u3.Titulo, u3.Artista)
		saved, _ := m.GetAlbum(email)
		_ = m.UpdateAlbum(saved)
		err = m.DeleteAlbum(email)
		assert.Nil(t, err)
	})

}
