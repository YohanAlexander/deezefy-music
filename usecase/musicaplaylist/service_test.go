package musicaplaylist

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicaplaylist"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/musicaplaylist"

	"github.com/stretchr/testify/assert"
)

func newFixtureMusicaPlaylist() *der.MusicaPlaylist {
	return &der.MusicaPlaylist{
		Musica:   1,
		Playlist: "Indie Rock",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureMusicaPlaylist()
		_, _, err := m.CreateMusicaPlaylist(u.Musica, u.Playlist)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureMusicaPlaylist()
	u2 := newFixtureMusicaPlaylist()
	u2.Musica = 2
	u2.Playlist = "Pop Rock"

	musica, playlist, _ := m.CreateMusicaPlaylist(u1.Musica, u1.Playlist)
	_, _, _ = m.CreateMusicaPlaylist(u2.Musica, u2.Playlist)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchMusicaPlaylists("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchMusicaPlaylists("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListMusicaPlaylists()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetMusicaPlaylist(musica, playlist)
		assert.Nil(t, err)
	})

	t.Run("get by Musica", func(t *testing.T) {
		_, err := m.GetMusicaPlaylistByMusica(musica)
		assert.NotNil(t, err)
	})

	t.Run("get by Playlist", func(t *testing.T) {
		_, err := m.GetMusicaPlaylistByPlaylist(playlist)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureMusicaPlaylist()
		musica, playlist, err := m.CreateMusicaPlaylist(u.Musica, u.Playlist)
		assert.Nil(t, err)
		saved, _ := m.GetMusicaPlaylist(musica, playlist)
		assert.Nil(t, m.UpdateMusicaPlaylist(saved))
		_, err = m.GetMusicaPlaylist(musica, playlist)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureMusicaPlaylist()
	u2 := newFixtureMusicaPlaylist()
	u2.Musica = 2
	u2.Playlist = "Pop Rock"
	musica, Playlist, _ := m.CreateMusicaPlaylist(u2.Musica, u2.Playlist)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteMusicaPlaylist(u1.Musica, u1.Playlist)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteMusicaPlaylist(musica, Playlist)
		assert.Nil(t, err)
		_, err = m.GetMusicaPlaylist(musica, Playlist)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureMusicaPlaylist()
		musica, playlist, _ := m.CreateMusicaPlaylist(u3.Musica, u3.Playlist)
		saved, _ := m.GetMusicaPlaylist(musica, playlist)
		_ = m.UpdateMusicaPlaylist(saved)
		err = m.DeleteMusicaPlaylist(musica, playlist)
		assert.Nil(t, err)
	})

}
