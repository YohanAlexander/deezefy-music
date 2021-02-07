package perfil

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/stretchr/testify/assert"
)

func newFixturePerfil() *entity.Perfil {
	return &entity.Perfil{
		ID: 815,
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email:    "someone@deezefy.com",
				Password: "12345678",
				Birthday: "1998-05-27",
			},
			PrimeiroNome: "Imagine",
			Sobrenome:    "Dragons",
		},
		InformacoesRelevantes: "Mais ouvidas",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		p := newFixturePerfil()
		_, err := m.CreatePerfil(p.Ouvinte.Usuario.Email, p.Ouvinte.Usuario.Password, p.Ouvinte.Usuario.Birthday, p.Ouvinte.PrimeiroNome, p.Ouvinte.Sobrenome, p.InformacoesRelevantes, p.ID)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	p1 := newFixturePerfil()
	p2 := newFixturePerfil()
	p2.ID = 200
	p2.Ouvinte.Usuario.Email = "someone2@deezefy.com"

	email, _ := m.CreatePerfil(p1.Ouvinte.Usuario.Email, p1.Ouvinte.Usuario.Password, p1.Ouvinte.Usuario.Birthday, p1.Ouvinte.PrimeiroNome, p1.Ouvinte.Sobrenome, p1.InformacoesRelevantes, p1.ID)
	_, _ = m.CreatePerfil(p2.Ouvinte.Usuario.Email, p2.Ouvinte.Usuario.Password, p2.Ouvinte.Usuario.Birthday, p2.Ouvinte.PrimeiroNome, p2.Ouvinte.Sobrenome, p2.InformacoesRelevantes, p2.ID)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchPerfils("someone@deezefy.com")
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
		p := newFixturePerfil()
		email, err := m.CreatePerfil(p.Ouvinte.Usuario.Email, p.Ouvinte.Usuario.Password, p.Ouvinte.Usuario.Birthday, p.Ouvinte.PrimeiroNome, p.Ouvinte.Sobrenome, p.InformacoesRelevantes, p.ID)
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
	p1 := newFixturePerfil()
	p2 := newFixturePerfil()
	p2.ID = 200
	p2.Ouvinte.Usuario.Email = "someone2@deezefy.com"
	email, _ := m.CreatePerfil(p2.Ouvinte.Usuario.Email, p2.Ouvinte.Usuario.Password, p2.Ouvinte.Usuario.Birthday, p2.Ouvinte.PrimeiroNome, p2.Ouvinte.Sobrenome, p2.InformacoesRelevantes, p2.ID)

	t.Run("delete", func(t *testing.T) {

		err := m.DeletePerfil(p1.Ouvinte.Usuario.Email)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeletePerfil(email)
		assert.Nil(t, err)
		_, err = m.GetPerfil(email)
		assert.Equal(t, entity.ErrNotFound, err)

		p3 := newFixturePerfil()
		email, _ := m.CreatePerfil(p3.Ouvinte.Usuario.Email, p3.Ouvinte.Usuario.Password, p3.Ouvinte.Usuario.Birthday, p3.Ouvinte.PrimeiroNome, p3.Ouvinte.Sobrenome, p3.InformacoesRelevantes, p3.ID)
		saved, _ := m.GetPerfil(email)
		_ = m.UpdatePerfil(saved)
		err = m.DeletePerfil(email)
		assert.Nil(t, err)
	})

}
