package ouvinteseguirartista

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
	ouvinte "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
	ouvinteseguirartista "github.com/yohanalexander/deezefy-music/usecase/relationship/ouvinteseguirartista/mock"
)

func Test_seguir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Artista := artista.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSeguirArtista := ouvinteseguirartista.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSeguirArtistaHandlers(r, *n, Ouvinte, Artista, OuvinteSeguirArtista)
	path, err := r.GetRoute("seguir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/seguir/{artista_email}", path)
	handler := seguir(Ouvinte, Artista, OuvinteSeguirArtista)
	r.Handle("/v1/{ouvinte_email}/seguir/{artista_email}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/seguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/seguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		OuvinteSeguirArtista.EXPECT().Seguir(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/seguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desseguir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Artista := artista.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSeguirArtista := ouvinteseguirartista.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSeguirArtistaHandlers(r, *n, Ouvinte, Artista, OuvinteSeguirArtista)
	path, err := r.GetRoute("desseguir").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/desseguir/{artista_email}", path)
	handler := desseguir(Ouvinte, Artista, OuvinteSeguirArtista)
	r.Handle("/v1/{ouvinte_email}/desseguir/{artista_email}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desseguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desseguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		OuvinteSeguirArtista.EXPECT().Desseguir(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desseguir/%s", ts.URL, o.Usuario.Email, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
