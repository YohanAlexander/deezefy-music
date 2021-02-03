package musicagenero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicagenero"

	"github.com/stretchr/testify/assert"
)

func newFixtureMusicaGenero() *der.MusicaGenero {
	return &der.MusicaGenero{
		Musica: 1,
		Genero: "Indie Rock",
	}
}

func TestCreate(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureMusicaGenero()
		_, _, err := m.CreateMusicaGenero(u.Musica, u.Genero)
		assert.Nil(t, err)
	})

}

func TestSearchAndFind(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureMusicaGenero()
	u2 := newFixtureMusicaGenero()
	u2.Musica = 2
	u2.Genero = "Pop Rock"

	musica, genero, _ := m.CreateMusicaGenero(u1.Musica, u1.Genero)
	_, _, _ = m.CreateMusicaGenero(u2.Musica, u2.Genero)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchMusicaGeneros("Indie Rock")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))

		c, err = m.SearchMusicaGeneros("nonedio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListMusicaGeneros()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		_, err := m.GetMusicaGenero(musica, genero)
		assert.Nil(t, err)
	})

	t.Run("get by Musica", func(t *testing.T) {
		_, err := m.GetMusicaGeneroByMusica(musica)
		assert.NotNil(t, err)
	})

	t.Run("get by genero", func(t *testing.T) {
		_, err := m.GetMusicaGeneroByGenero(genero)
		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update", func(t *testing.T) {
		repo := newInmem()
		m := NewService(repo)
		u := newFixtureMusicaGenero()
		musica, genero, err := m.CreateMusicaGenero(u.Musica, u.Genero)
		assert.Nil(t, err)
		saved, _ := m.GetMusicaGenero(musica, genero)
		assert.Nil(t, m.UpdateMusicaGenero(saved))
		_, err = m.GetMusicaGenero(musica, genero)
		assert.Nil(t, err)
	})

}

func TestDelete(t *testing.T) {

	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureMusicaGenero()
	u2 := newFixtureMusicaGenero()
	u2.Musica = 2
	u2.Genero = "genero2@email.com"
	musica, genero, _ := m.CreateMusicaGenero(u2.Musica, u2.Genero)

	t.Run("delete", func(t *testing.T) {

		err := m.DeleteMusicaGenero(u1.Musica, u1.Genero)
		assert.Equal(t, entity.ErrNotFound, err)

		err = m.DeleteMusicaGenero(musica, genero)
		assert.Nil(t, err)
		_, err = m.GetMusicaGenero(musica, genero)
		assert.Equal(t, entity.ErrNotFound, err)

		u3 := newFixtureMusicaGenero()
		musica, genero, _ := m.CreateMusicaGenero(u3.Musica, u3.Genero)
		saved, _ := m.GetMusicaGenero(musica, genero)
		_ = m.UpdateMusicaGenero(saved)
		err = m.DeleteMusicaGenero(musica, genero)
		assert.Nil(t, err)
	})

}
