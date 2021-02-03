package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlbum(t *testing.T) {

	t.Run("Album criado com sucesso", func(t *testing.T) {
		a, err := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		assert.Nil(t, err)
		assert.Equal(t, a.Titulo, "Viva la vida")
	})

}

func TestAlbum_Validate(t *testing.T) {

	type artista struct {
		email         string
		password      string
		birthday      string
		nomeartistico string
		biografia     string
		anoformacao   int
	}

	type test struct {
		name          string
		artista       artista
		id            int
		titulo        string
		anolancamento int
		want          error
	}

	tests := []test{
		{
			name: "Campos válidos",
			artista: artista{
				email:         "vancejoy@gmail.com",
				password:      "new_password",
				birthday:      "2006-01-02",
				nomeartistico: "Vance Joy",
				biografia:     "Australian Singer",
				anoformacao:   2006,
			},
			id:            1,
			titulo:        "Yellow",
			anolancamento: 2000,
			want:          nil,
		},
		{
			name: "ID inválido",
			artista: artista{
				email:         "vancejoy@gmail.com",
				password:      "new_password",
				birthday:      "2006-01-02",
				nomeartistico: "Vance Joy",
				biografia:     "Australian Singer",
				anoformacao:   2006,
			},
			id:            0,
			titulo:        "Yellow",
			anolancamento: 2000,
			want:          ErrInvalidEntity,
		},
		{
			name: "Título inválido",
			artista: artista{
				email:         "vancejoy@gmail.com",
				password:      "new_password",
				birthday:      "2006-01-02",
				nomeartistico: "Vance Joy",
				biografia:     "Australian Singer",
				anoformacao:   2006,
			},
			id:            1,
			titulo:        "",
			anolancamento: 2000,
			want:          ErrInvalidEntity,
		},
		{
			name: "AnoLancamento inválido",
			artista: artista{
				email:         "vancejoy@gmail.com",
				password:      "new_password",
				birthday:      "2006-01-02",
				nomeartistico: "Vance Joy",
				biografia:     "Australian Singer",
				anoformacao:   2006,
			},
			id:            1,
			titulo:        "Yellow",
			anolancamento: 200,
			want:          ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewAlbum(tc.artista.email, tc.artista.password, tc.artista.birthday, tc.artista.nomeartistico, tc.artista.biografia, tc.titulo, tc.artista.anoformacao, tc.anolancamento, tc.id)
			assert.Equal(t, err, tc.want)
		})
	}

}

func TestAddSalvouAlbum(t *testing.T) {

	t.Run("Ouvinte criado com sucesso", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(a.Salvou))
	})

	t.Run("Ouvinte já registrado", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.AddOuvinte(*o)
		assert.Nil(t, err)
		o, _ = NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err = a.AddOuvinte(*o)
		assert.Equal(t, ErrOuvinteRegistered, err)
	})

}

func TestRemoveSalvouAlbum(t *testing.T) {

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		err := a.RemoveOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("Ouvinte removido com sucesso", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		err := a.RemoveOuvinte(*o)
		assert.Nil(t, err)
	})

}

func TestGetSalvouAlbum(t *testing.T) {

	t.Run("Ouvinte cadastrado encontrado", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_ = a.AddOuvinte(*o)
		ouvinte, err := a.GetOuvinte(*o)
		assert.Nil(t, err)
		assert.Equal(t, ouvinte, *o)
	})

	t.Run("Ouvinte não cadastrado", func(t *testing.T) {
		a, _ := NewAlbum("coldplay@gmail.com", "somepassword", "2018-02-10", "Coldplay", "British Band", "Viva la vida", 2000, 2006, 1)
		o, _ := NewOuvinte("vancejoy@gmail.com", "somepassword", "2018-02-10", "Vance", "Joy")
		_, err := a.GetOuvinte(*o)
		assert.Equal(t, ErrNotFound, err)
	})

}
