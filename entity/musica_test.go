package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMusica(t *testing.T) {

	t.Run("Musica criada com sucesso", func(t *testing.T) {
		m, err := NewMusica(1, 420, "Creep")
		assert.Nil(t, err)
		assert.Equal(t, m.Nome, "Creep")
	})

}

func TestMusica_Validate(t *testing.T) {

	type test struct {
		name    string
		id      int
		duracao int
		nome    string
		want    error
	}

	tests := []test{
		{
			name:    "Campos válidos",
			id:      1,
			duracao: 100,
			nome:    "Creep",
			want:    nil,
		},
		{
			name:    "ID inválido",
			id:      0,
			duracao: 100,
			nome:    "Creep",
			want:    ErrInvalidEntity,
		},
		{
			name:    "Nome inválido",
			id:      1,
			duracao: 100,
			nome:    "",
			want:    ErrInvalidEntity,
		},
		{
			name:    "Duração inválida",
			id:      1,
			duracao: 10,
			nome:    "Creep",
			want:    ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMusica(tc.id, tc.duracao, tc.nome)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddCurtiuMusica(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := m.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(m.Curtiu))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := m.AddOuvinte(*o)
		assert.Nil(t, err)
		o, _ = NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err = m.AddOuvinte(*o)
		assert.Equal(t, ErrOuvinteRegistered, err)
	})

}

func TestRemoveCurtiuMusica(t *testing.T) {

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := m.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = m.AddOuvinte(*o)
		err := m.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetCurtiuMusica(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = m.AddOuvinte(*o)
		ouvinte, err := m.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica(1, 420, "Creep")
		_, err := m.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddGravaArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := m.AddArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(m.Gravou))
	})

	t.Run("Artista já registrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := m.AddArtista(*a)
		assert.Nil(t, err)
		a, _ = NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err = m.AddArtista(*a)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveGravaArtista(t *testing.T) {

	t.Run("Artista não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := m.RemoveArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Artista removido com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = m.AddArtista(*a)
		err := m.RemoveArtista(*a)
		assert.Nil(t, err)
	})

}

func TestGetGravaArtista(t *testing.T) {

	t.Run("Artista cadastrado encontrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = m.AddArtista(*a)
		artista, err := m.GetArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, artista, *a)
	})

	t.Run("Artista não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_, err := m.GetArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}
