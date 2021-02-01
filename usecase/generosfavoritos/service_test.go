package generosfavoritos

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/generosfavoritos"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/generosfavoritos"

	"github.com/stretchr/testify/assert"
)

func newFixtureGenerosFavoritos() *der.GenerosFavoritos {
	return &der.GenerosFavoritos{
		Perfil:  10,
		Genero:  "Indie Rock",
		Ouvinte: "ouvinte@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureGenerosFavoritos()
		_, _, _, err := m.CreateGenerosFavoritos(u.Perfil, u.Genero, u.Ouvinte)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureGenerosFavoritos()
	u2 := newFixtureGenerosFavoritos()
	u2.Perfil = 20
	u2.Genero = "Pop Rock"
	u2.Ouvinte = "ouvinte2@email.com"

	perfil, genero, ouvinte, _ := m.CreateGenerosFavoritos(u1.Perfil, u1.Genero, u1.Ouvinte)
	_, _, _, _ = m.CreateGenerosFavoritos(u2.Perfil, u2.Genero, u2.Ouvinte)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchGenerosFavoritos("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchGenerosFavoritos("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListGenerosFavoritos()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetGenerosFavoritos(perfil, genero, ouvinte)
		assert.Nil(t, err)
	})

	t.Run("get by Perfil", func(t *testing.T) {
		_, err := m.GetGenerosFavoritosByPerfil(perfil)
		assert.NotNil(t, err)
	})

	t.Run("get by Genero", func(t *testing.T) {
		_, err := m.GetGenerosFavoritosByGenero(genero)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureGenerosFavoritos()
		perfil, genero, ouvinte, err := m.CreateGenerosFavoritos(u.Perfil, u.Genero, u.Ouvinte)
		assert.Nil(t, err)
		saved, _ := m.GetGenerosFavoritos(perfil, genero, ouvinte)
		assert.Nil(t, m.UpdateGenerosFavoritos(saved))
		_, err = m.GetGenerosFavoritos(perfil, genero, ouvinte)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureGenerosFavoritos()
	u2 := newFixtureGenerosFavoritos()
	u2.Perfil = 20
	u2.Ouvinte = "ouvinte2@email.com"
	perfil, genero, ouvinte, _ := m.CreateGenerosFavoritos(u2.Perfil, u2.Genero, u2.Ouvinte)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteGenerosFavoritos(u1.Perfil, u1.Genero, u1.Ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteGenerosFavoritos(perfil, genero, ouvinte)
		assert.Nil(t, err)
		_, err = m.GetGenerosFavoritos(perfil, genero, ouvinte)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureGenerosFavoritos()
		perfil, genero, ouvinte, _ := m.CreateGenerosFavoritos(u3.Perfil, u3.Genero, u3.Ouvinte)
		saved, _ := m.GetGenerosFavoritos(perfil, genero, ouvinte)
		_ = m.UpdateGenerosFavoritos(saved)
		err = m.DeleteGenerosFavoritos(perfil, genero, ouvinte)
		assert.Nil(t, err)
	})

}
