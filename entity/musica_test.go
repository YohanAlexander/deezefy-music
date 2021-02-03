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

func TestAddMusicaPlaylist(t *testing.T) {

	t.Run("Playlist criado com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		err := m.AddPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(m.Playlists))
	})

	t.Run("Playlist já registrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		err := m.AddPlaylist(*p)
		assert.Nil(t, err)
		p, _ = NewPlaylist("Indie Rock", "ativo")
		err = m.AddPlaylist(*p)
		assert.Equal(t, ErrPlaylistRegistered, err)
	})

}

func TestRemoveMusicaPlaylist(t *testing.T) {

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		err := m.RemovePlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Playlist removido com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		_ = m.AddPlaylist(*p)
		err := m.RemovePlaylist(*p)
		assert.Nil(t, err)
	})

}

func TestGetMusicaPlaylist(t *testing.T) {

	t.Run("Playlist cadastrado encontrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		_ = m.AddPlaylist(*p)
		playlist, err := m.GetPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, playlist, *p)
	})

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		p, _ := NewPlaylist("Indie Rock", "ativo")
		_, err := m.GetPlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddMusicaAlbum(t *testing.T) {

	t.Run("Album criado com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := m.AddAlbum(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(m.Albums))
	})

	t.Run("Album já registrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := m.AddAlbum(*a)
		assert.Nil(t, err)
		a, _ = NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err = m.AddAlbum(*a)
		assert.Equal(t, ErrAlbumRegistered, err)
	})

}

func TestRemoveMusicaAlbum(t *testing.T) {

	t.Run("Album não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := m.RemoveAlbum(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Album removido com sucesso", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = m.AddAlbum(*a)
		err := m.RemoveAlbum(*a)
		assert.Nil(t, err)
	})

}

func TestGetMusicaAlbum(t *testing.T) {

	t.Run("Album cadastrado encontrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = m.AddAlbum(*a)
		album, err := m.GetAlbum(*a)
		assert.Nil(t, err)
		assert.Equal(t, album, *a)
	})

	t.Run("Album não cadastrado", func(t *testing.T) {
		m, _ := NewMusica(1, 420, "Creep")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_, err := m.GetAlbum(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}
