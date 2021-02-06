package artistapossuirgenero

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	artista "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	genero "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
	artistapossuirgenero "github.com/yohanalexander/deezefy-music/usecase/relationship/artistapossuirgenero/mock"
)

func Test_possuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Artista := artista.NewMockUseCase(controller)
	ArtistaPossuirGenero := artistapossuirgenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaPossuirGeneroHandlers(r, *n, Artista, Genero, ArtistaPossuirGenero)
	path, err := r.GetRoute("possuir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{artista_email}/possuir/{genero_nome}", path)
	handler := possuir(Artista, Genero, ArtistaPossuirGenero)
	r.Handle("/v1/{artista_email}/possuir/{genero_nome}", handler)
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/possuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/possuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(m, nil)
		ArtistaPossuirGenero.EXPECT().Possuir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/possuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_despossuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Artista := artista.NewMockUseCase(controller)
	ArtistaPossuirGenero := artistapossuirgenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaPossuirGeneroHandlers(r, *n, Artista, Genero, ArtistaPossuirGenero)
	path, err := r.GetRoute("despossuir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{artista_email}/despossuir/{genero_nome}", path)
	handler := despossuir(Artista, Genero, ArtistaPossuirGenero)
	r.Handle("/v1/{artista_email}/despossuir/{genero_nome}", handler)
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/despossuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/despossuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(m, nil)
		ArtistaPossuirGenero.EXPECT().Despossuir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/despossuir/%s", ts.URL, o.Usuario.Email, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
