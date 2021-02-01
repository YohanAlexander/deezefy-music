package salvaalbum

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaalbum"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/salvaalbum"

	"github.com/stretchr/testify/assert"
)

func newFixtureSalvaAlbum() *der.SalvaAlbum {
	return &der.SalvaAlbum{
		Album:   10,
		Ouvinte: "ouvinte@email.com",
		Artista: "artista@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSalvaAlbum()
		_, _, _, err := m.CreateSalvaAlbum(u.Album, u.Ouvinte, u.Artista)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSalvaAlbum()
	u2 := newFixtureSalvaAlbum()
	u2.Album = 20
	u2.Ouvinte = "ouvinte2@email.com"

	album, ouvinte, artista, _ := m.CreateSalvaAlbum(u1.Album, u1.Ouvinte, u1.Artista)
	_, _, _, _ = m.CreateSalvaAlbum(u2.Album, u2.Ouvinte, u2.Artista)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchSalvaAlbums("ouvinte@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchSalvaAlbums("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListSalvaAlbums()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetSalvaAlbum(album, ouvinte, artista)
		assert.Nil(t, err)
	})

	t.Run("get by Album", func(t *testing.T) {
		_, err := m.GetSalvaAlbumByAlbum(album)
		assert.NotNil(t, err)
	})

	t.Run("get by ouvinte", func(t *testing.T) {
		_, err := m.GetSalvaAlbumByOuvinte(ouvinte)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSalvaAlbum()
		album, ouvinte, artista, err := m.CreateSalvaAlbum(u.Album, u.Ouvinte, u.Artista)
		assert.Nil(t, err)
		saved, _ := m.GetSalvaAlbum(album, ouvinte, artista)
		assert.Nil(t, m.UpdateSalvaAlbum(saved))
		_, err = m.GetSalvaAlbum(album, ouvinte, artista)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSalvaAlbum()
	u2 := newFixtureSalvaAlbum()
	u2.Album = 20
	u2.Ouvinte = "ouvinte2@email.com"
	album, ouvinte, artista, _ := m.CreateSalvaAlbum(u2.Album, u2.Ouvinte, u2.Artista)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteSalvaAlbum(u1.Album, u1.Ouvinte, u1.Artista)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteSalvaAlbum(album, ouvinte, artista)
		assert.Nil(t, err)
		_, err = m.GetSalvaAlbum(album, ouvinte, artista)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureSalvaAlbum()
		album, ouvinte, artista, _ := m.CreateSalvaAlbum(u3.Album, u3.Ouvinte, u3.Artista)
		saved, _ := m.GetSalvaAlbum(album, ouvinte, artista)
		_ = m.UpdateSalvaAlbum(saved)
		err = m.DeleteSalvaAlbum(album, ouvinte, artista)
		assert.Nil(t, err)
	})

}
