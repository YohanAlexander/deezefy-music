package ouvintecurtirmusica

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
	musica "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	ouvinte "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
	ouvintecurtirmusica "github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintecurtirmusica/mock"
)

func Test_curtir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteCurtirMusica := ouvintecurtirmusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteCurtirMusicaHandlers(r, *n, Ouvinte, Musica, OuvinteCurtirMusica)
	path, err := r.GetRoute("curtir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/curtir/{musica_id}", path)
	handler := curtir(Ouvinte, Musica, OuvinteCurtirMusica)
	r.Handle("/v1/{ouvinte_email}/curtir/{musica_id}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/curtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/curtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		OuvinteCurtirMusica.EXPECT().Curtir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/curtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_descurtir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteCurtirMusica := ouvintecurtirmusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteCurtirMusicaHandlers(r, *n, Ouvinte, Musica, OuvinteCurtirMusica)
	path, err := r.GetRoute("descurtir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/descurtir/{musica_id}", path)
	handler := descurtir(Ouvinte, Musica, OuvinteCurtirMusica)
	r.Handle("/v1/{ouvinte_email}/descurtir/{musica_id}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/descurtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/descurtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		OuvinteCurtirMusica.EXPECT().Descurtir(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/descurtir/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
