package playlist

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixturePlaylist() *entity.Playlist {
	return &entity.Playlist{
		Nome:   "Indie Rock",
		Status: "ativo",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixturePlaylist()
		_, err := m.CreatePlaylist(u.Nome, u.Status)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePlaylist()
	u2 := newFixturePlaylist()
	u2.Nome = "Pop Rock"

	email, _ := m.CreatePlaylist(u1.Nome, u1.Status)
	_, _ = m.CreatePlaylist(u2.Nome, u2.Status)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchPlaylists("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchPlaylists("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPlaylists()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetPlaylist(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixturePlaylist()
		email, err := m.CreatePlaylist(u.Nome, u.Status)
		assert.Nil(t, err)
		saved, _ := m.GetPlaylist(email)
		assert.Nil(t, m.UpdatePlaylist(saved))
		_, err = m.GetPlaylist(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePlaylist()
	u2 := newFixturePlaylist()
	u2.Nome = "someone2@deezefy.com"
	email, _ := m.CreatePlaylist(u2.Nome, u2.Status)

	t.Run("delete", func(t *testing.T) {

		err := m.DeletePlaylist(u1.Nome)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeletePlaylist(email)
		assert.Nil(t, err)
		_, err = m.GetPlaylist(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixturePlaylist()
		email, _ := m.CreatePlaylist(u3.Nome, u3.Status)
		saved, _ := m.GetPlaylist(email)
		_ = m.UpdatePlaylist(saved)
		err = m.DeletePlaylist(email)
		assert.Nil(t, err)
	})

}
