package grava

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/grava"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/grava"

	"github.com/stretchr/testify/assert"
)

func newFixtureGrava() *der.Grava {
	return &der.Grava{
		Musica:  1,
		Artista: "artista@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureGrava()
		_, _, err := m.CreateGrava(u.Musica, u.Artista)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureGrava()
	u2 := newFixtureGrava()
	u2.Musica = 2
	u2.Artista = "artista2@email.com"

	musica, artista, _ := m.CreateGrava(u1.Musica, u1.Artista)
	_, _, _ = m.CreateGrava(u2.Musica, u2.Artista)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchGravas("Artista@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchGravas("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListGravas()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetGrava(musica, artista)
		assert.Nil(t, err)
	})

	t.Run("get by Musica", func(t *testing.T) {
		_, err := m.GetGravaByMusica(musica)
		assert.NotNil(t, err)
	})

	t.Run("get by Artista", func(t *testing.T) {
		_, err := m.GetGravaByArtista(artista)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureGrava()
		musica, artista, err := m.CreateGrava(u.Musica, u.Artista)
		assert.Nil(t, err)
		saved, _ := m.GetGrava(musica, artista)
		assert.Nil(t, m.UpdateGrava(saved))
		_, err = m.GetGrava(musica, artista)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureGrava()
	u2 := newFixtureGrava()
	u2.Musica = 2
	u2.Artista = "Artista2@email.com"
	musica, Artista, _ := m.CreateGrava(u2.Musica, u2.Artista)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteGrava(u1.Musica, u1.Artista)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteGrava(musica, Artista)
		assert.Nil(t, err)
		_, err = m.GetGrava(musica, Artista)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureGrava()
		musica, Artista, _ := m.CreateGrava(u3.Musica, u3.Artista)
		saved, _ := m.GetGrava(musica, Artista)
		_ = m.UpdateGrava(saved)
		err = m.DeleteGrava(musica, Artista)
		assert.Nil(t, err)
	})

}
