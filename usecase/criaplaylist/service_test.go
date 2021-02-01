package criaplaylist

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/criaplaylist"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/criaplaylist"

	"github.com/stretchr/testify/assert"
)

func newFixtureCriaPlaylist() *der.CriaPlaylist {
	return &der.CriaPlaylist{
		DataCriacao: "2010-05-15",
		Playlist:    "Indie Rock",
		Usuario:     "usuario@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureCriaPlaylist()
		_, _, err := m.CreateCriaPlaylist(u.DataCriacao, u.Playlist, u.Usuario)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureCriaPlaylist()
	u2 := newFixtureCriaPlaylist()
	u2.Playlist = "Pop Rock"
	u2.Usuario = "usuario2@email.com"

	playlist, usuario, _ := m.CreateCriaPlaylist(u1.DataCriacao, u1.Playlist, u1.Usuario)
	_, _, _ = m.CreateCriaPlaylist(u2.DataCriacao, u2.Playlist, u2.Usuario)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchCriaPlaylists("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchCriaPlaylists("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListCriaPlaylists()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetCriaPlaylist(playlist, usuario)
		assert.Nil(t, err)
	})

	t.Run("get by Playlist", func(t *testing.T) {
		_, err := m.GetCriaPlaylistByPlaylist(playlist)
		assert.NotNil(t, err)
	})

	t.Run("get by Usuario", func(t *testing.T) {
		_, err := m.GetCriaPlaylistByUsuario(usuario)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureCriaPlaylist()
		playlist, usuario, err := m.CreateCriaPlaylist(u.DataCriacao, u.Playlist, u.Usuario)
		assert.Nil(t, err)
		saved, _ := m.GetCriaPlaylist(playlist, usuario)
		assert.Nil(t, m.UpdateCriaPlaylist(saved))
		_, err = m.GetCriaPlaylist(playlist, usuario)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureCriaPlaylist()
	u2 := newFixtureCriaPlaylist()
	u2.Playlist = "Pop Rock"
	u2.Usuario = "usuario2@email.com"
	playlist, usuario, _ := m.CreateCriaPlaylist(u2.DataCriacao, u2.Playlist, u2.Usuario)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteCriaPlaylist(u1.Playlist, u1.Usuario)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteCriaPlaylist(playlist, usuario)
		assert.Nil(t, err)
		_, err = m.GetCriaPlaylist(playlist, usuario)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureCriaPlaylist()
		playlist, usuario, _ := m.CreateCriaPlaylist(u3.DataCriacao, u3.Playlist, u3.Usuario)
		saved, _ := m.GetCriaPlaylist(playlist, usuario)
		_ = m.UpdateCriaPlaylist(saved)
		err = m.DeleteCriaPlaylist(playlist, usuario)
		assert.Nil(t, err)
	})

}
