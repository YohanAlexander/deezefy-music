package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlaylist(t *testing.T) {

	t.Run("Playlist criada com sucesso", func(t *testing.T) {
		p, err := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		assert.Nil(t, err)
		assert.Equal(t, p.Status, "ativo")
	})

}

func TestPlaylist_Validate(t *testing.T) {

	type test struct {
		name        string
		nome        string
		status      string
		datacriacao string
		want        error
	}

	tests := []test{
		{
			name:        "Campos válidos",
			nome:        "Indie Rock",
			status:      "ativo",
			datacriacao: "2006-01-02",
			want:        nil,
		},
		{
			name:        "Nome inválido",
			nome:        "",
			status:      "ativo",
			datacriacao: "2006-01-02",
			want:        ErrInvalidEntity,
		},
		{
			name:        "Status inválido",
			nome:        "Indie Rock",
			status:      "ativa",
			datacriacao: "2006-01-02",
			want:        ErrInvalidEntity,
		},
		{
			name:        "DataCriacao inválida",
			nome:        "Indie Rock",
			status:      "ativo",
			datacriacao: "2006/01/02",
			want:        ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPlaylist(tc.nome, tc.status, tc.datacriacao)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddSalvouPlaylist(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := p.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p.Salvou))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
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
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := p.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = p.AddOuvinte(*o)
		err := p.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetSalvouPlaylist(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = p.AddOuvinte(*o)
		ouvinte, err := p.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		_, err := p.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddPlaylistMusica(t *testing.T) {

	t.Run("Musica criado com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		err := p.AddMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p.Musicas))
	})

	t.Run("Musica já registrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		err := p.AddMusica(*m)
		assert.Nil(t, err)
		m, _ = NewMusica("Creep", 420, 1)
		err = p.AddMusica(*m)
		assert.Equal(t, ErrMusicaRegistered, err)
	})

}

func TestRemovePlaylistMusica(t *testing.T) {

	t.Run("Musica não cadastrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		err := p.RemoveMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Musica removido com sucesso", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		_ = p.AddMusica(*m)
		err := p.RemoveMusica(*m)
		assert.Nil(t, err)
	})

}

func TestGetPlaylistMusica(t *testing.T) {

	t.Run("Musica cadastrado encontrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		_ = p.AddMusica(*m)
		musica, err := p.GetMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, musica, *m)
	})

	t.Run("Musica não cadastrado", func(t *testing.T) {
		p, _ := NewPlaylist("Indie Rock", "ativo", "2006-01-02")
		m, _ := NewMusica("Creep", 420, 1)
		_, err := p.GetMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

}
