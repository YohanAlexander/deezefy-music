package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenero(t *testing.T) {

	t.Run("Genero criado com sucesso", func(t *testing.T) {
		g, err := NewGenero("Indie Rock", "rock")
		assert.Nil(t, err)
		assert.Equal(t, g.Estilo, "rock")
	})

}

func TestGenero_Validate(t *testing.T) {

	type test struct {
		name   string
		nome   string
		estilo string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			nome:   "Indie Rock",
			estilo: "rock",
			want:   nil,
		},
		{
			name:   "Nome inválido",
			nome:   "",
			estilo: "rock",
			want:   ErrInvalidEntity,
		},
		{
			name:   "Estilo inválido",
			nome:   "Indie Rock",
			estilo: "indie",
			want:   ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGenero(tc.nome, tc.estilo)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddGeneroArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := g.AddArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(g.Artistas))
	})

	t.Run("Artista já registrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := g.AddArtista(*a)
		assert.Nil(t, err)
		a, _ = NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err = g.AddArtista(*a)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveGeneroArtista(t *testing.T) {

	t.Run("Artista não cadastrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := g.RemoveArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Artista removido com sucesso", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = g.AddArtista(*a)
		err := g.RemoveArtista(*a)
		assert.Nil(t, err)
	})

}

func TestGetGeneroArtista(t *testing.T) {

	t.Run("Artista cadastrado encontrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = g.AddArtista(*a)
		artista, err := g.GetArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, artista, *a)
	})

	t.Run("Artista não cadastrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_, err := g.GetArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddGeneroMusica(t *testing.T) {

	t.Run("Musica criado com sucesso", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		err := g.AddMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(g.Musicas))
	})

	t.Run("Musica já registrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		err := g.AddMusica(*m)
		assert.Nil(t, err)
		m, _ = NewMusica(1, 420, "Creep")
		err = g.AddMusica(*m)
		assert.Equal(t, ErrMusicaRegistered, err)
	})

}

func TestRemoveGeneroMusica(t *testing.T) {

	t.Run("Musica não cadastrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		err := g.RemoveMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Musica removido com sucesso", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		_ = g.AddMusica(*m)
		err := g.RemoveMusica(*m)
		assert.Nil(t, err)
	})

}

func TestGetGeneroMusica(t *testing.T) {

	t.Run("Musica cadastrado encontrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		_ = g.AddMusica(*m)
		Musica, err := g.GetMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, Musica, *m)
	})

	t.Run("Musica não cadastrado", func(t *testing.T) {
		g, _ := NewGenero("Indie Rock", "rock")
		m, _ := NewMusica(1, 420, "Creep")
		_, err := g.GetMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

}
