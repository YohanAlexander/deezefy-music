package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerfil(t *testing.T) {

	t.Run("Perfil criado com sucesso", func(t *testing.T) {
		p, err := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		assert.Nil(t, err)
		assert.Equal(t, p.InformacoesRelevantes, "Where is my love")
	})

}

func TestPerfil_Validate(t *testing.T) {

	type ouvinte struct {
		email        string
		password     string
		birthday     string
		primeironome string
		sobrenome    string
	}

	type test struct {
		name                  string
		ouvinte               ouvinte
		id                    int
		informacoesrelevantes string
		want                  error
	}

	tests := []test{
		{
			name:                  "Campos válidos",
			informacoesrelevantes: "Where is my love",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   1,
			want: nil,
		},
		{
			name:                  "InformaçõesRelevantes inválidas",
			informacoesrelevantes: "",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   1,
			want: ErrInvalidEntity,
		},
		{
			name:                  "ID inválido",
			informacoesrelevantes: "Where is my love",
			ouvinte: ouvinte{
				email:        "vancejoy@gmail.com",
				password:     "new_password",
				birthday:     "2006-01-02",
				primeironome: "Vance",
				sobrenome:    "Joy",
			},
			id:   0,
			want: ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPerfil(tc.ouvinte.email, tc.ouvinte.password, tc.ouvinte.birthday, tc.ouvinte.primeironome, tc.ouvinte.sobrenome, tc.informacoesrelevantes, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddArtistasFavoritos(t *testing.T) {

	t.Run("Artista criado com sucesso", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := p.AddArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p.ArtistasFavoritos))
	})

	t.Run("Artista já registrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := p.AddArtista(*a)
		assert.Nil(t, err)
		a, _ = NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err = p.AddArtista(*a)
		assert.Equal(t, ErrArtistaRegistered, err)
	})

}

func TestRemoveArtistasFavoritos(t *testing.T) {

	t.Run("Artista não cadastrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		err := p.RemoveArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Artista removido com sucesso", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = p.AddArtista(*a)
		err := p.RemoveArtista(*a)
		assert.Nil(t, err)
	})

}

func TestGetArtistasFavoritos(t *testing.T) {

	t.Run("Artista cadastrado encontrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_ = p.AddArtista(*a)
		artista, err := p.GetArtista(*a)
		assert.Nil(t, err)
		assert.Equal(t, artista, *a)
	})

	t.Run("Artista não cadastrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		a, _ := NewArtista("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance Joy", "Australian Singer", 2006)
		_, err := p.GetArtista(*a)
		assert.Equal(t, ErrNotFound, err)
	})

}

func TestAddPerfilGenero(t *testing.T) {

	t.Run("Genero criado com sucesso", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		err := p.AddGenero(*g)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p.GenerosFavoritos))
	})

	t.Run("Genero já registrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		err := p.AddGenero(*g)
		assert.Nil(t, err)
		g, _ = NewGenero("Indie Rock", "rock")
		err = p.AddGenero(*g)
		assert.Equal(t, ErrGeneroRegistered, err)
	})

}

func TestRemovePerfilGenero(t *testing.T) {

	t.Run("Genero não cadastrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		err := p.RemoveGenero(*g)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Genero removido com sucesso", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		_ = p.AddGenero(*g)
		err := p.RemoveGenero(*g)
		assert.Nil(t, err)
	})

}

func TestGetPerfilGenero(t *testing.T) {

	t.Run("Genero cadastrado encontrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		_ = p.AddGenero(*g)
		genero, err := p.GetGenero(*g)
		assert.Nil(t, err)
		assert.Equal(t, genero, *g)
	})

	t.Run("Genero não cadastrado", func(t *testing.T) {
		p, _ := NewPerfil("syml@spotify.com", "somepassword", "2018-02-10", "Vance", "Joy", "Where is my love", 1)
		g, _ := NewGenero("Indie Rock", "rock")
		_, err := p.GetGenero(*g)
		assert.Equal(t, ErrNotFound, err)
	})

}
