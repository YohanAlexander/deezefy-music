package perfil

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixturePerfil() *entity.Perfil {
	return &entity.Perfil{
		ID:                    815,
		Ouvinte:               "someone@spotify.com",
		InformacoesRelevantes: "Mais ouvidas",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixturePerfil()
		_, err := m.CreatePerfil(u.Ouvinte, u.InformacoesRelevantes, u.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePerfil()
	u2 := newFixturePerfil()
	u2.ID = 200
	u2.InformacoesRelevantes = "Mais compartilhadas"

	email, _ := m.CreatePerfil(u1.Ouvinte, u1.InformacoesRelevantes, u1.ID)
	_, _ = m.CreatePerfil(u2.Ouvinte, u2.InformacoesRelevantes, u2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchPerfils("Mais ouvidas")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchPerfils("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPerfils()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetPerfil(email)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixturePerfil()
		email, err := m.CreatePerfil(u.Ouvinte, u.InformacoesRelevantes, u.ID)
		assert.Nil(t, err)
		saved, _ := m.GetPerfil(email)
		assert.Nil(t, m.UpdatePerfil(saved))
		_, err = m.GetPerfil(email)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePerfil()
	u2 := newFixturePerfil()
	u2.ID = 200
	email, _ := m.CreatePerfil(u2.Ouvinte, u2.InformacoesRelevantes, u2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeletePerfil(u1.ID)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeletePerfil(email)
		assert.Nil(t, err)
		_, err = m.GetPerfil(email)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixturePerfil()
		email, _ := m.CreatePerfil(u3.Ouvinte, u3.InformacoesRelevantes, u3.ID)
		saved, _ := m.GetPerfil(email)
		_ = m.UpdatePerfil(saved)
		err = m.DeletePerfil(email)
		assert.Nil(t, err)
	})

}
