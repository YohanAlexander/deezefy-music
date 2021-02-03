package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlaylist(t *testing.T) {

	t.Run("Playlist criada com sucesso", func(t *testing.T) {
		p, err := NewPlaylist("Indie Rock", "ativo")
		assert.Nil(t, err)
		assert.Equal(t, p.Status, "ativo")
	})

}

func TestPlaylist_Validate(t *testing.T) {

	type test struct {
		name   string
		nome   string
		status string
		want   error
	}

	tests := []test{
		{
			name:   "Campos válidos",
			nome:   "Indie Rock",
			status: "ativo",
			want:   nil,
		},
		{
			name:   "Nome inválido",
			nome:   "",
			status: "ativo",
			want:   ErrInvalidEntity,
		},
		{
			name:   "Status inválido",
			nome:   "Indie Rock",
			status: "ativa",
			want:   ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPlaylist(tc.nome, tc.status)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddSalvouPlaylist(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := p.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p.Salvou))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := p.AddOuvinte(*o)
		assert.Nil(t, err)
		o, _ = NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err = p.AddOuvinte(*o)
		assert.Equal(t, ErrOuvinteRegistered, err)
	})

}

func TestRemoveSalvouPlaylist(t *testing.T) {

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := p.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = p.AddOuvinte(*o)
		err := p.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetSalvouPlaylist(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = p.AddOuvinte(*o)
		ouvinte, err := p.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		_, err := p.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}
