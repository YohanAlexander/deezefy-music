package ouvintesalvaralbum

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
	album "github.com/yohanalexander/deezefy-music/usecase/entity/album/mock"
	ouvinte "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
	ouvintesalvaralbum "github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvaralbum/mock"
)

func Test_salvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Album := album.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSalvarAlbum := ouvintesalvaralbum.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSalvarAlbumHandlers(r, *n, Ouvinte, Album, OuvinteSalvarAlbum)
	path, err := r.GetRoute("salvar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/salvar/{album_id}", path)
	handler := salvar(Ouvinte, Album, OuvinteSalvarAlbum)
	r.Handle("/v1/{ouvinte_email}/salvar/{album_id}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Album.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Album.EXPECT().GetAlbum(a.ID).Return(a, nil)
		OuvinteSalvarAlbum.EXPECT().Salvar(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_dessalvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Album := album.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSalvarAlbum := ouvintesalvaralbum.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSalvarAlbumHandlers(r, *n, Ouvinte, Album, OuvinteSalvarAlbum)
	path, err := r.GetRoute("dessalvar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/dessalvar/{album_id}", path)
	handler := dessalvar(Ouvinte, Album, OuvinteSalvarAlbum)
	r.Handle("/v1/{ouvinte_email}/dessalvar/{album_id}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Album.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Album.EXPECT().GetAlbum(a.ID).Return(a, nil)
		OuvinteSalvarAlbum.EXPECT().Dessalvar(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%d", ts.URL, o.Usuario.Email, a.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
