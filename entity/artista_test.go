package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtista(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		a, err := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		assert.Nil(t, err)
		assert.Equal(t, a.NomeArtistico, "Vance Joy")
	})

}

func TestArtista_Validate(t *testing.T) {

	type user struct {
		email    string
		password string
		birthday string
	}

	type test struct {
		name          string
		usuario       user
		nomeartistico string
		biografia     string
		anoformacao   int
		want          error
	}

	tests := []test{
		{
			name: "Campos válidos",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          nil,
		},
		{
			name: "NomeArtistico inválido",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "",
			biografia:     "Australian Singer",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "Biografia inválida",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "",
			anoformacao:   2006,
			want:          ErrInvalidEntity,
		},
		{
			name: "AnoFormacao inválido (2000)",
			usuario: user{
				email:    "vancejoy@gmail.com",
				password: "new_password",
				birthday: "2006-01-02",
			},
			nomeartistico: "Vance Joy",
			biografia:     "Australian Singer",
			anoformacao:   98,
			want:          ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArtista(tc.usuario.email, tc.usuario.password, tc.usuario.birthday, tc.nomeartistico, tc.biografia, tc.anoformacao)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddOuvinte(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Seguidores))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		o, _ = NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err = a.AddOuvinte(*o)
		assert.Equal(t, ErrOuvinteRegistered, err)
	})

}

func TestRemoveOuvinte(t *testing.T) {

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		err := a.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetOuvinte(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		ouvinte, err := a.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_, err := a.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddGravaMusica(t *testing.T) {

	t.Run("Musica criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.AddMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Grava))
	})

	t.Run("Musica já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.AddMusica(*m)
		assert.Nil(t, err)
		m, _ = NewMusica(1, 420, "Creep")
		err = a.AddMusica(*m)
		assert.Equal(t, ErrMusicaRegistered, err)
	})

}

func TestRemoveGravaMusica(t *testing.T) {

	t.Run("Musica não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		err := a.RemoveMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Musica removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_ = a.AddMusica(*m)
		err := a.RemoveMusica(*m)
		assert.Nil(t, err)
	})

}

func TestGetGravaMusica(t *testing.T) {

	t.Run("Musica cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_ = a.AddMusica(*m)
		musica, err := a.GetMusica(*m)
		assert.Nil(t, err)
		assert.Equal(t, musica, *m)
	})

	t.Run("Musica não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		m, _ := NewMusica(1, 420, "Creep")
		_, err := a.GetMusica(*m)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddPerfil(t *testing.T) {

	t.Run("Perfil criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		err := a.AddPerfil(*p)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Perfis))
	})

	t.Run("Perfil já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		err := a.AddPerfil(*p)
		assert.Nil(t, err)
		p, _ = NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		err = a.AddPerfil(*p)
		assert.Equal(t, ErrPerfilRegistered, err)
	})

}

func TestRemovePerfil(t *testing.T) {

	t.Run("Perfil não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		err := a.RemovePerfil(*p)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Perfil removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		_ = a.AddPerfil(*p)
		err := a.RemovePerfil(*p)
		assert.Nil(t, err)
	})

}

func TestGetPerfil(t *testing.T) {

	t.Run("Perfil cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		_ = a.AddPerfil(*p)
		perfil, err := a.GetPerfil(*p)
		assert.Nil(t, err)
		assert.Equal(t, perfil, *p)
	})

	t.Run("Perfil não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		_, err := a.GetPerfil(*p)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddArtistaGenero(t *testing.T) {

	t.Run("Genero criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		err := a.AddGenero(*g)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Generos))
	})

	t.Run("Genero já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		err := a.AddGenero(*g)
		assert.Nil(t, err)
		g, _ = NewGenero("Indie Rock", "rock")
		err = a.AddGenero(*g)
		assert.Equal(t, ErrGeneroRegistered, err)
	})

}

func TestRemoveArtistaGenero(t *testing.T) {

	t.Run("Genero não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		err := a.RemoveGenero(*g)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Genero removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		_ = a.AddGenero(*g)
		err := a.RemoveGenero(*g)
		assert.Nil(t, err)
	})

}

func TestGetArtistaGenero(t *testing.T) {

	t.Run("Genero cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		_ = a.AddGenero(*g)
		genero, err := a.GetGenero(*g)
		assert.Nil(t, err)
		assert.Equal(t, genero, *g)
	})

	t.Run("Genero não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		g, _ := NewGenero("Indie Rock", "rock")
		_, err := a.GetGenero(*g)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddAlbumArtista(t *testing.T) {

	t.Run("Album criado com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := a.AddAlbum(*b)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Albums))
	})

	t.Run("Album já registrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := a.AddAlbum(*b)
		assert.Nil(t, err)
		b, _ = NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err = a.AddAlbum(*b)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveAlbumArtista(t *testing.T) {

	t.Run("Album não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		err := a.RemoveAlbum(*b)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Album removido com sucesso", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = a.AddAlbum(*b)
		err := a.RemoveAlbum(*b)
		assert.Nil(t, err)
	})

}

func TestGetAlbumArtista(t *testing.T) {

	t.Run("Album cadastrado encontrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_ = a.AddAlbum(*b)
		artista, err := a.GetAlbum(*b)
		assert.Nil(t, err)
		assert.Equal(t, artista, *b)
	})

	t.Run("Album não cadastrado", func(t *testing.T) {
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		b, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		_, err := a.GetAlbum(*b)
		assert.Equal(t, ErrNotFound, err)
	})

}
