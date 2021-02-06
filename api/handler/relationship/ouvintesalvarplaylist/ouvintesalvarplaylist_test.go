package ouvintesalvarplaylist

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
	ouvinte "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
	playlist "github.com/yohanalexander/deezefy-music/usecase/entity/playlist/mock"
	ouvintesalvarplaylist "github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvarplaylist/mock"
)

func Test_salvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Playlist := playlist.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSalvarPlaylist := ouvintesalvarplaylist.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSalvarPlaylistHandlers(r, *n, Ouvinte, Playlist, OuvinteSalvarPlaylist)
	path, err := r.GetRoute("salvar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/salvar/{playlist_nome}", path)
	handler := salvar(Ouvinte, Playlist, OuvinteSalvarPlaylist)
	r.Handle("/v1/{ouvinte_email}/salvar/{playlist_nome}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Playlist.EXPECT().GetPlaylist(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Playlist.EXPECT().GetPlaylist(a.Nome).Return(a, nil)
		OuvinteSalvarPlaylist.EXPECT().Salvar(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/salvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_dessalvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Playlist := playlist.NewMockUseCase(controller)
	Ouvinte := ouvinte.NewMockUseCase(controller)
	OuvinteSalvarPlaylist := ouvintesalvarplaylist.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteSalvarPlaylistHandlers(r, *n, Ouvinte, Playlist, OuvinteSalvarPlaylist)
	path, err := r.GetRoute("dessalvar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{ouvinte_email}/dessalvar/{playlist_nome}", path)
	handler := dessalvar(Ouvinte, Playlist, OuvinteSalvarPlaylist)
	r.Handle("/v1/{ouvinte_email}/dessalvar/{playlist_nome}", handler)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Playlist.EXPECT().GetPlaylist(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Playlist{
			Nome: "Playlist",
		}
		Ouvinte.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		Playlist.EXPECT().GetPlaylist(a.Nome).Return(a, nil)
		OuvinteSalvarPlaylist.EXPECT().Dessalvar(o, a).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/dessalvar/%s", ts.URL, o.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
