package artistasfavoritos

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistasfavoritos"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/artistasfavoritos"

	"github.com/stretchr/testify/assert"
)

func newFixtureArtistasFavoritos() *der.ArtistasFavoritos {
	return &der.ArtistasFavoritos{
		Perfil:  10,
		Ouvinte: "ouvinte@email.com",
		Artista: "artista@email.com",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtistasFavoritos()
		_, _, _, err := m.CreateArtistasFavoritos(u.Perfil, u.Ouvinte, u.Artista)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtistasFavoritos()
	u2 := newFixtureArtistasFavoritos()
	u2.Perfil = 20
	u2.Ouvinte = "ouvinte2@email.com"

	perfil, ouvinte, artista, _ := m.CreateArtistasFavoritos(u1.Perfil, u1.Ouvinte, u1.Artista)
	_, _, _, _ = m.CreateArtistasFavoritos(u2.Perfil, u2.Ouvinte, u2.Artista)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchArtistasFavoritoss("ouvinte@email.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchArtistasFavoritoss("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListArtistasFavoritoss()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetArtistasFavoritos(perfil, ouvinte, artista)
		assert.Nil(t, err)
	})

	t.Run("get by Perfil", func(t *testing.T) {
		_, err := m.GetArtistasFavoritosByPerfil(perfil)
		assert.NotNil(t, err)
	})

	t.Run("get by Ouvinte", func(t *testing.T) {
		_, err := m.GetArtistasFavoritosByOuvinte(ouvinte)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureArtistasFavoritos()
		perfil, ouvinte, artista, err := m.CreateArtistasFavoritos(u.Perfil, u.Ouvinte, u.Artista)
		assert.Nil(t, err)
		saved, _ := m.GetArtistasFavoritos(perfil, ouvinte, artista)
		assert.Nil(t, m.UpdateArtistasFavoritos(saved))
		_, err = m.GetArtistasFavoritos(perfil, ouvinte, artista)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureArtistasFavoritos()
	u2 := newFixtureArtistasFavoritos()
	u2.Perfil = 20
	u2.Ouvinte = "ouvinte2@email.com"
	perfil, ouvinte, artista, _ := m.CreateArtistasFavoritos(u2.Perfil, u2.Ouvinte, u2.Artista)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteArtistasFavoritos(u1.Perfil, u1.Ouvinte, u1.Artista)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteArtistasFavoritos(perfil, ouvinte, artista)
		assert.Nil(t, err)
		_, err = m.GetArtistasFavoritos(perfil, ouvinte, artista)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureArtistasFavoritos()
		perfil, ouvinte, artista, _ := m.CreateArtistasFavoritos(u3.Perfil, u3.Ouvinte, u3.Artista)
		saved, _ := m.GetArtistasFavoritos(perfil, ouvinte, artista)
		_ = m.UpdateArtistasFavoritos(saved)
		err = m.DeleteArtistasFavoritos(perfil, ouvinte, artista)
		assert.Nil(t, err)
	})

}
