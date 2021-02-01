package artistagenero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistagenero"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/artistagenero"

	"github.com/stretchr/testify/assert"
)

func newFixtureArtistaGenero() *der.ArtistaGenero {
	return &der.ArtistaGenero{
		Artista: "artista@email.com",
		Genero:  "Indie Rock",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtistaGenero()
		_, _, err := m.CreateArtistaGenero(u.Artista, u.Genero)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtistaGenero()
	u2 := newFixtureArtistaGenero()
	u2.Artista = "artista2@email.com"
	u2.Genero = "Genero2@email.com"

	artista, genero, _ := m.CreateArtistaGenero(u1.Artista, u1.Genero)
	_, _, _ = m.CreateArtistaGenero(u2.Artista, u2.Genero)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchArtistaGeneros("artista@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchArtistaGeneros("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListArtistaGeneros()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetArtistaGenero(artista, genero)
		assert.Nil(t, err)
	})

	t.Run("get by Artista", func(t *testing.T) {
		_, err := m.GetArtistaGeneroByArtista(artista)
		assert.NotNil(t, err)
	})

	t.Run("get by Genero", func(t *testing.T) {
		_, err := m.GetArtistaGeneroByGenero(genero)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtistaGenero()
		artista, genero, err := m.CreateArtistaGenero(u.Artista, u.Genero)
		assert.Nil(t, err)
		saved, _ := m.GetArtistaGenero(artista, genero)
		assert.Nil(t, m.UpdateArtistaGenero(saved))
		_, err = m.GetArtistaGenero(artista, genero)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtistaGenero()
	u2 := newFixtureArtistaGenero()
	u2.Artista = "artista2@email.com"
	u2.Genero = "Genero2@email.com"
	artista, genero, _ := m.CreateArtistaGenero(u2.Artista, u2.Genero)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteArtistaGenero(u1.Artista, u1.Genero)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteArtistaGenero(artista, genero)
		assert.Nil(t, err)
		_, err = m.GetArtistaGenero(artista, genero)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureArtistaGenero()
		artista, genero, _ := m.CreateArtistaGenero(u3.Artista, u3.Genero)
		saved, _ := m.GetArtistaGenero(artista, genero)
		_ = m.UpdateArtistaGenero(saved)
		err = m.DeleteArtistaGenero(artista, genero)
		assert.Nil(t, err)
	})

}
