package albummusica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/albummusica"

	"github.com/stretchr/testify/assert"
)

func newFixtureAlbumMusica() *der.AlbumMusica {
	return &der.AlbumMusica{
		Album:   10,
		Musica:  10,
		Artista: "artista@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureAlbumMusica()
		_, _, _, err := m.CreateAlbumMusica(u.Album, u.Musica, u.Artista)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureAlbumMusica()
	u2 := newFixtureAlbumMusica()
	u2.Album = 20
	u2.Musica = 20
	u2.Artista = "artista2@email.com"

	album, musica, artista, _ := m.CreateAlbumMusica(u1.Album, u1.Musica, u1.Artista)
	_, _, _, _ = m.CreateAlbumMusica(u2.Album, u2.Musica, u2.Artista)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchAlbumMusicas("artista@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchAlbumMusicas("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListAlbumMusicas()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetAlbumMusica(album, musica, artista)
		assert.Nil(t, err)
	})

	t.Run("get by Album", func(t *testing.T) {
		_, err := m.GetAlbumMusicaByAlbum(album)
		assert.NotNil(t, err)
	})

	t.Run("get by Musica", func(t *testing.T) {
		_, err := m.GetAlbumMusicaByMusica(musica)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureAlbumMusica()
		album, musica, artista, err := m.CreateAlbumMusica(u.Album, u.Musica, u.Artista)
		assert.Nil(t, err)
		saved, _ := m.GetAlbumMusica(album, musica, artista)
		assert.Nil(t, m.UpdateAlbumMusica(saved))
		_, err = m.GetAlbumMusica(album, musica, artista)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureAlbumMusica()
	u2 := newFixtureAlbumMusica()
	u2.Album = 20
	u2.Musica = 20
	album, musica, artista, _ := m.CreateAlbumMusica(u2.Album, u2.Musica, u2.Artista)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteAlbumMusica(u1.Album, u1.Musica, u1.Artista)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteAlbumMusica(album, musica, artista)
		assert.Nil(t, err)
		_, err = m.GetAlbumMusica(album, musica, artista)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureAlbumMusica()
		album, musica, artista, _ := m.CreateAlbumMusica(u3.Album, u3.Musica, u3.Artista)
		saved, _ := m.GetAlbumMusica(album, musica, artista)
		_ = m.UpdateAlbumMusica(saved)
		err = m.DeleteAlbumMusica(album, musica, artista)
		assert.Nil(t, err)
	})

}
