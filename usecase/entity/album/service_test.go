package album

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureAlbum() *entity.Album {
	return &entity.Album{
		ID:            815,
		AnoLancamento: 1998,
		Titulo:        "Cage The Elephant",
		Artista: entity.Artista{
			Usuario: entity.Usuario{
				Email:    "someone@deezefy.com",
				Password: "12345678",
				Birthday: "1998-05-27",
			},
			NomeArtistico: "Imagine Dragons",
			Biografia:     "Indie Rock Band",
			AnoFormacao:   1998,
		},
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		a := newFixtureAlbum()
		_, err := m.CreateAlbum(a.Artista.Usuario.Email, a.Artista.Usuario.Password, a.Artista.Usuario.Birthday, a.Artista.NomeArtistico, a.Artista.Biografia, a.Titulo, a.Artista.AnoFormacao, a.AnoLancamento, a.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	a1 := newFixtureAlbum()
	a2 := newFixtureAlbum()
	a2.ID = 200
	a2.Titulo = "Radioactive"

	email, _ := m.CreateAlbum(a1.Artista.Usuario.Email, a1.Artista.Usuario.Password, a1.Artista.Usuario.Birthday, a1.Artista.NomeArtistico, a1.Artista.Biografia, a1.Titulo, a1.Artista.AnoFormacao, a1.AnoLancamento, a1.ID)
	_, _ = m.CreateAlbum(a2.Artista.Usuario.Email, a2.Artista.Usuario.Password, a2.Artista.Usuario.Birthday, a2.Artista.NomeArtistico, a2.Artista.Biografia, a2.Titulo, a2.Artista.AnoFormacao, a2.AnoLancamento, a2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchAlbums("Cage The Elephant")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchAlbums("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListAlbums()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetAlbum(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		a := newFixtureAlbum()
		email, err := m.CreateAlbum(a.Artista.Usuario.Email, a.Artista.Usuario.Password, a.Artista.Usuario.Birthday, a.Artista.NomeArtistico, a.Artista.Biografia, a.Titulo, a.Artista.AnoFormacao, a.AnoLancamento, a.ID)
		assert.Nil(t, err)
		saved, _ := m.GetAlbum(email)
		assert.Nil(t, m.UpdateAlbum(saved))
		_, err = m.GetAlbum(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	a1 := newFixtureAlbum()
	a2 := newFixtureAlbum()
	a2.ID = 200
	email, _ := m.CreateAlbum(a2.Artista.Usuario.Email, a2.Artista.Usuario.Password, a2.Artista.Usuario.Birthday, a2.Artista.NomeArtistico, a2.Artista.Biografia, a2.Titulo, a2.Artista.AnoFormacao, a2.AnoLancamento, a2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteAlbum(a1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteAlbum(email)
		assert.Nil(t, err)
		_, err = m.GetAlbum(email)
		assert.Equal(t, entity.ErrNotFound, err)

		a3 := newFixtureAlbum()
		email, _ := m.CreateAlbum(a3.Artista.Usuario.Email, a3.Artista.Usuario.Password, a3.Artista.Usuario.Birthday, a3.Artista.NomeArtistico, a3.Artista.Biografia, a3.Titulo, a3.Artista.AnoFormacao, a3.AnoLancamento, a3.ID)
		saved, _ := m.GetAlbum(email)
		_ = m.UpdateAlbum(saved)
		err = m.DeleteAlbum(email)
		assert.Nil(t, err)
	})

}
