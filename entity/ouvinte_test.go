package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOuvinte(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		o, err := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		assert.Nil(t, err)
		assert.Equal(t, o.PrimeiroNome, "Vance")
	})

}

func TestOuvinte_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name         string
		usuario      user
		primeironome string
		sobrenome    string
		want         error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "Vance",
			sobrenome:    "Joy",
			want:         nil,
		},
		{
			name: "PrimeiroNome inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "",
			sobrenome:    "Joy",
			want:         ErrInvalidEntity,
		},
		{
			name: "Sobrenome inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			primeironome: "Vance",
			sobrenome:    "",
			want:         ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewOuvinte(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.primeironome, tc.sobrenome)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddCriaPlaylist(t *testing.T) {

	t.Run("Playlist criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.AddCriaPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Cria))
	})

	t.Run("Playlist já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.AddCriaPlaylist(*p)
		assert.Nil(t, err)
		p, _ = NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err = o.AddCriaPlaylist(*p)
		assert.Equal(t, ErrPlaylistRegistered, err)
	})

}

func TestRemoveCriaPlaylist(t *testing.T) {

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.RemoveCriaPlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Playlist removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_ = o.AddCriaPlaylist(*p)
		err := o.RemoveCriaPlaylist(*p)
		assert.Nil(t, err)
	})

}

func TestGetCriaPlaylist(t *testing.T) {

	t.Run("Playlist cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_ = o.AddCriaPlaylist(*p)
		playlist, err := o.GetCriaPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, playlist, *p)
	})

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_, err := o.GetCriaPlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddTelefone(t *testing.T) {

	t.Run("Telefone criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Telefones))
	})

	t.Run("Telefone já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.AddTelefone("+5579999999999")
		assert.Nil(t, err)
		err = o.AddTelefone("+5579999999999")
		assert.Equal(t, ErrPhoneRegistered, err)
	})

}

func TestRemoveTelefone(t *testing.T) {

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := o.RemoveTelefone("+5579999999999")
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Telefone removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		err := o.RemoveTelefone("+5579999999999")
		assert.Nil(t, err)
	})

}

func TestGetTelefone(t *testing.T) {

	t.Run("Telefone cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = o.AddTelefone("+5579999999999")
		telefone, err := o.GetTelefone("+5579999999999")
		assert.Nil(t, err)
		assert.Equal(t, telefone, "+5579999999999")
	})

	t.Run("Telefone não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_, err := o.GetTelefone("+5579999999999")
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.AddArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Seguindo))
	})

	t.Run("Artista já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.AddArtista(*a)
		assert.Nil(t, err)
		a, _ = NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err = o.AddArtista(*a)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveArtista(t *testing.T) {

	t.Run("Artista não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := o.RemoveArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Artista removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = o.AddArtista(*a)
		err := o.RemoveArtista(*a)
		assert.Nil(t, err)
	})

}

func TestGetArtista(t *testing.T) {

	t.Run("Artista cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = o.AddArtista(*a)
		artista, err := o.GetArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, artista, *a)
	})

	t.Run("Artista não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_, err := o.GetArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddCurteMusica(t *testing.T) {

	t.Run("Musica criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		err := o.AddMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Curtidas))
	})

	t.Run("Musica já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		err := o.AddMusica(*m)
		assert.Nil(t, err)
		m, _ = NewMusica("Creep", 420, 1)
		err = o.AddMusica(*m)
		assert.Equal(t, ErrMusicaRegistered, err)
	})

}

func TestRemoveCurteMusica(t *testing.T) {

	t.Run("Musica não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		err := o.RemoveMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Musica removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		_ = o.AddMusica(*m)
		err := o.RemoveMusica(*m)
		assert.Nil(t, err)
	})

}

func TestGetCurteMusica(t *testing.T) {

	t.Run("Musica cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		_ = o.AddMusica(*m)
		musica, err := o.GetMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, musica, *m)
	})

	t.Run("Musica não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		m, _ := NewMusica("Creep", 420, 1)
		_, err := o.GetMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddPlaylist(t *testing.T) {

	t.Run("Playlist criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.AddPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Playlists))
	})

	t.Run("Playlist já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.AddPlaylist(*p)
		assert.Nil(t, err)
		p, _ = NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err = o.AddPlaylist(*p)
		assert.Equal(t, ErrPlaylistRegistered, err)
	})

}

func TestRemovePlaylist(t *testing.T) {

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		err := o.RemovePlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Playlist removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_ = o.AddPlaylist(*p)
		err := o.RemovePlaylist(*p)
		assert.Nil(t, err)
	})

}

func TestGetPlaylist(t *testing.T) {

	t.Run("Playlist cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_ = o.AddPlaylist(*p)
		playlist, err := o.GetPlaylist(*p)
		assert.Nil(t, err)
		assert.Equal(t, playlist, *p)
	})

	t.Run("Playlist não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		p, _ := NewPlaylist("steve.jobs@apple.com", "new_password", "2006-01-02", "Indie Rock", "ativo", "2006-01-02")
		_, err := o.GetPlaylist(*p)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddAlbum(t *testing.T) {

	t.Run("Album criado com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := o.AddAlbum(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(o.Albums))
	})

	t.Run("Album já registrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := o.AddAlbum(*a)
		assert.Nil(t, err)
		a, _ = NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err = o.AddAlbum(*a)
		assert.Equal(t, ErrAlbumRegistered, err)
	})

}

func TestRemoveAlbum(t *testing.T) {

	t.Run("Album não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := o.RemoveAlbum(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Album removido com sucesso", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = o.AddAlbum(*a)
		err := o.RemoveAlbum(*a)
		assert.Nil(t, err)
	})

}

func TestGetAlbum(t *testing.T) {

	t.Run("Album cadastrado encontrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = o.AddAlbum(*a)
		Album, err := o.GetAlbum(*a)
		assert.Nil(t, err)
		assert.Equal(t, Album, *a)
	})

	t.Run("Album não cadastrado", func(t *testing.T) {
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_, err := o.GetAlbum(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}
