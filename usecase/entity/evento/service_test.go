package evento

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureEvento() *entity.Evento {
	return &entity.Evento{
		ID:   815,
		Nome: "Lollapalooza",
		Data: "1998-05-27",
		Usuario: entity.Usuario{
			Email:    "someone@deezefy.com",
			Password: "12345678",
			Birthday: "1998-05-27",
		},
		Local: entity.Local{
			Cidade: "São Paulo",
			Pais:   "Brazil",
			ID:     1,
		},
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		e := newFixtureEvento()
		_, err := m.CreateEvento(e.Usuario.Email, e.Usuario.Password, e.Usuario.Birthday, e.Nome, e.Data, e.Local.Cidade, e.Local.Pais, e.Local.ID, e.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	e1 := newFixtureEvento()
	e2 := newFixtureEvento()
	e2.ID = 200
	e2.Nome = "Rock in Rio"

	email, _ := m.CreateEvento(e1.Usuario.Email, e1.Usuario.Password, e1.Usuario.Birthday, e1.Nome, e1.Data, e1.Local.Cidade, e1.Local.Pais, e1.Local.ID, e1.ID)
	_, _ = m.CreateEvento(e2.Usuario.Email, e2.Usuario.Password, e2.Usuario.Birthday, e2.Nome, e2.Data, e2.Local.Cidade, e2.Local.Pais, e2.Local.ID, e2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchEventos("Lollapalooza")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchEventos("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListEventos()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetEvento(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		e := newFixtureEvento()
		email, err := m.CreateEvento(e.Usuario.Email, e.Usuario.Password, e.Usuario.Birthday, e.Nome, e.Data, e.Local.Cidade, e.Local.Pais, e.Local.ID, e.ID)
		assert.Nil(t, err)
		saved, _ := m.GetEvento(email)
		assert.Nil(t, m.UpdateEvento(saved))
		_, err = m.GetEvento(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	e1 := newFixtureEvento()
	e2 := newFixtureEvento()
	e2.ID = 200
	email, _ := m.CreateEvento(e2.Usuario.Email, e2.Usuario.Password, e2.Usuario.Birthday, e2.Nome, e2.Data, e2.Local.Cidade, e2.Local.Pais, e2.Local.ID, e2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteEvento(e1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteEvento(email)
		assert.Nil(t, err)
		_, err = m.GetEvento(email)
		assert.Equal(t, entity.ErrNotFound, err)

		e3 := newFixtureEvento()
		email, _ := m.CreateEvento(e3.Usuario.Email, e3.Usuario.Password, e3.Usuario.Birthday, e3.Nome, e3.Data, e3.Local.Cidade, e3.Local.Pais, e3.Local.ID, e3.ID)
		saved, _ := m.GetEvento(email)
		_ = m.UpdateEvento(saved)
		err = m.DeleteEvento(email)
		assert.Nil(t, err)
	})

}
