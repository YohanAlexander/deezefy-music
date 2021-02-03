package ocorre

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ocorre"

	"github.com/stretchr/testify/assert"
)

func newFixtureOcorre() *der.Ocorre {
	return &der.Ocorre{
		Data:    "2010-08-21",
		Artista: "artista@email.com",
		Usuario: "artista@email.com",
		Local:   10,
		Evento:  10,
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureOcorre()
		_, _, _, _, err := m.CreateOcorre(u.Data, u.Artista, u.Usuario, u.Local, u.Evento)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureOcorre()
	u2 := newFixtureOcorre()
	u2.Local = 20
	u2.Evento = 20
	u2.Artista = "artista2@email.com"

	artista, usuario, local, evento, _ := m.CreateOcorre(u1.Data, u1.Artista, u1.Usuario, u1.Local, u1.Evento)
	_, _, _, _, _ = m.CreateOcorre(u2.Data, u2.Artista, u2.Usuario, u2.Local, u2.Evento)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchOcorres("artista@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchOcorres("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListOcorres()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetOcorre(artista, usuario, local, evento)
		assert.Nil(t, err)
	})

	t.Run("get by Local", func(t *testing.T) {
		_, err := m.GetOcorreByLocal(local)
		assert.NotNil(t, err)
	})

	t.Run("get by Evento", func(t *testing.T) {
		_, err := m.GetOcorreByEvento(evento)
		assert.NotNil(t, err)
	})

	t.Run("get by Artista", func(t *testing.T) {
		_, err := m.GetOcorreByArtista(artista)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureOcorre()
		artista, usuario, local, evento, err := m.CreateOcorre(u.Data, u.Artista, u.Usuario, u.Local, u.Evento)
		assert.Nil(t, err)
		saved, _ := m.GetOcorre(artista, usuario, local, evento)
		assert.Nil(t, m.UpdateOcorre(saved))
		_, err = m.GetOcorre(artista, usuario, local, evento)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureOcorre()
	u2 := newFixtureOcorre()
	u2.Local = 20
	u2.Evento = 20
	u2.Artista = "artista2@email.com"
	artista, usuario, local, evento, _ := m.CreateOcorre(u2.Data, u2.Artista, u2.Usuario, u2.Local, u2.Evento)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteOcorre(u1.Artista, u1.Usuario, u1.Local, u1.Evento)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteOcorre(artista, usuario, local, evento)
		assert.Nil(t, err)
		_, err = m.GetOcorre(artista, usuario, local, evento)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureOcorre()
		artista, usuario, local, evento, _ := m.CreateOcorre(u3.Data, u3.Artista, u3.Usuario, u3.Local, u3.Evento)
		saved, _ := m.GetOcorre(artista, usuario, local, evento)
		_ = m.UpdateOcorre(saved)
		err = m.DeleteOcorre(artista, usuario, local, evento)
		assert.Nil(t, err)
	})

}
