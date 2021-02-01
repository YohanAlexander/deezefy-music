package local

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/local"
	inmem "github.com/yohanalexander/deezefy-music/infrastructure/inmem/repository/local"

	"github.com/stretchr/testify/assert"
)

func newFixtureLocal() *der.Local {
	return &der.Local{
		ID:     815,
		Cidade: "Londres",
		Pais:   "Inglaterra",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureLocal()
		_, err := m.CreateLocal(u.Cidade, u.Pais, u.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureLocal()
	u2 := newFixtureLocal()
	u2.ID = 200
	u2.Cidade = "Cardiff"

	email, _ := m.CreateLocal(u1.Cidade, u1.Pais, u1.ID)
	_, _ = m.CreateLocal(u2.Cidade, u2.Pais, u2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchLocals("Londres")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchLocals("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListLocals()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetLocal(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := inmem.NewInmem()
		m := NewService(repo)
		u := newFixtureLocal()
		email, err := m.CreateLocal(u.Cidade, u.Pais, u.ID)
		assert.Nil(t, err)
		saved, _ := m.GetLocal(email)
		assert.Nil(t, m.UpdateLocal(saved))
		_, err = m.GetLocal(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := inmem.NewInmem()
	m := NewService(repo)
	u1 := newFixtureLocal()
	u2 := newFixtureLocal()
	u2.ID = 200
	email, _ := m.CreateLocal(u2.Cidade, u2.Pais, u2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteLocal(u1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteLocal(email)
		assert.Nil(t, err)
		_, err = m.GetLocal(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureLocal()
		email, _ := m.CreateLocal(u3.Cidade, u3.Pais, u3.ID)
		saved, _ := m.GetLocal(email)
		_ = m.UpdateLocal(saved)
		err = m.DeleteLocal(email)
		assert.Nil(t, err)
	})

}
