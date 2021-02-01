package salvaplaylist

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaplaylist"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/salvaplaylist"

	"github.com/stretchr/testify/assert"
)

func newFixtureSalvaPlaylist() *der.SalvaPlaylist {
	return &der.SalvaPlaylist{
		Playlist: "Indie Rock",
		Ouvinte:  "ouvinte@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSalvaPlaylist()
		_, _, err := m.CreateSalvaPlaylist(u.Playlist, u.Ouvinte)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSalvaPlaylist()
	u2 := newFixtureSalvaPlaylist()
	u2.Playlist = "Pop Rock"
	u2.Ouvinte = "ouvinte2@email.com"

	playlist, ouvinte, _ := m.CreateSalvaPlaylist(u1.Playlist, u1.Ouvinte)
	_, _, _ = m.CreateSalvaPlaylist(u2.Playlist, u2.Ouvinte)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchSalvaPlaylists("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchSalvaPlaylists("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListSalvaPlaylists()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetSalvaPlaylist(playlist, ouvinte)
		assert.Nil(t, err)
	})

	t.Run("get by Playlist", func(t *testing.T) {
		_, err := m.GetSalvaPlaylistByPlaylist(playlist)
		assert.NotNil(t, err)
	})

	t.Run("get by ouvinte", func(t *testing.T) {
		_, err := m.GetSalvaPlaylistByOuvinte(ouvinte)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureSalvaPlaylist()
		playlist, ouvinte, err := m.CreateSalvaPlaylist(u.Playlist, u.Ouvinte)
		assert.Nil(t, err)
		saved, _ := m.GetSalvaPlaylist(playlist, ouvinte)
		assert.Nil(t, m.UpdateSalvaPlaylist(saved))
		_, err = m.GetSalvaPlaylist(playlist, ouvinte)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureSalvaPlaylist()
	u2 := newFixtureSalvaPlaylist()
	u2.Playlist = "Playlist2@email.com"
	u2.Ouvinte = "ouvinte2@email.com"
	playlist, ouvinte, _ := m.CreateSalvaPlaylist(u2.Playlist, u2.Ouvinte)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteSalvaPlaylist(u1.Playlist, u1.Ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteSalvaPlaylist(playlist, ouvinte)
		assert.Nil(t, err)
		_, err = m.GetSalvaPlaylist(playlist, ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureSalvaPlaylist()
		playlist, ouvinte, _ := m.CreateSalvaPlaylist(u3.Playlist, u3.Ouvinte)
		saved, _ := m.GetSalvaPlaylist(playlist, ouvinte)
		_ = m.UpdateSalvaPlaylist(saved)
		err = m.DeleteSalvaPlaylist(playlist, ouvinte)
		assert.Nil(t, err)
	})

}
