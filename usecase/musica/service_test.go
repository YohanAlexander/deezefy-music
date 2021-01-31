package musica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musica"

	"github.com/stretchr/testify/assert"
)

func newFixtureMusica() *der.Musica {
	return &der.Musica{
		ID:      1998,
		Duracao: 815,
		Nome:    "Sultans of Swing",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureMusica()
		_, err := m.CreateMusica(u.ID, u.Duracao, u.Nome)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureMusica()
	u2 := newFixtureMusica()
	u2.ID = 200
	u2.Nome = "Radioactive"

	email, _ := m.CreateMusica(u1.ID, u1.Duracao, u1.Nome)
	_, _ = m.CreateMusica(u2.ID, u2.Duracao, u2.Nome)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchMusicas("Sultans of Swing")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchMusicas("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListMusicas()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetMusica(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureMusica()
		email, err := m.CreateMusica(u.ID, u.Duracao, u.Nome)
		assert.Nil(t, err)
		saved, _ := m.GetMusica(email)
		assert.Nil(t, m.UpdateMusica(saved))
		_, err = m.GetMusica(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureMusica()
	u2 := newFixtureMusica()
	u2.ID = 200
	email, _ := m.CreateMusica(u2.ID, u2.Duracao, u2.Nome)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteMusica(u1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteMusica(email)
		assert.Nil(t, err)
		_, err = m.GetMusica(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureMusica()
		email, _ := m.CreateMusica(u3.ID, u3.Duracao, u3.Nome)
		saved, _ := m.GetMusica(email)
		_ = m.UpdateMusica(saved)
		err = m.DeleteMusica(email)
		assert.Nil(t, err)
	})

}
