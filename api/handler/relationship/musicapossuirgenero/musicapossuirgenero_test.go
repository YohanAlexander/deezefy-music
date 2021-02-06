package musicapossuirgenero

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
	genero "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
	musica "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	musicapossuirgenero "github.com/yohanalexander/deezefy-music/usecase/relationship/musicapossuirgenero/mock"
)

func Test_possuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Musica := musica.NewMockUseCase(controller)
	MusicaPossuirGenero := musicapossuirgenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaPossuirGeneroHandlers(r, *n, Musica, Genero, MusicaPossuirGenero)
	path, err := r.GetRoute("possuir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{musica_id}/possuir/{genero_nome}", path)
	handler := possuir(Musica, Genero, MusicaPossuirGenero)
	r.Handle("/v1/{musica_id}/possuir/{genero_nome}", handler)
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/possuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/possuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(m, nil)
		MusicaPossuirGenero.EXPECT().Possuir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/possuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_despossuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Musica := musica.NewMockUseCase(controller)
	MusicaPossuirGenero := musicapossuirgenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaPossuirGeneroHandlers(r, *n, Musica, Genero, MusicaPossuirGenero)
	path, err := r.GetRoute("despossuir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{musica_id}/despossuir/{genero_nome}", path)
	handler := despossuir(Musica, Genero, MusicaPossuirGenero)
	r.Handle("/v1/{musica_id}/despossuir/{genero_nome}", handler)
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/despossuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/despossuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Musica{
			ID: 1,
		}
		m := &entity.Genero{
			Nome: "Genero",
		}
		Musica.EXPECT().GetMusica(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(m.Nome).Return(m, nil)
		MusicaPossuirGenero.EXPECT().Despossuir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/despossuir/%s", ts.URL, o.ID, m.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
