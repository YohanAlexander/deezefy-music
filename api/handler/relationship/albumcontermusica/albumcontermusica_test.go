package albumcontermusica

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
	musica "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	albumcontermusica "github.com/yohanalexander/deezefy-music/usecase/relationship/albumcontermusica/mock"
)

func Test_conter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Album := album.NewMockUseCase(controller)
	AlbumConterMusica := albumcontermusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumConterMusicaHandlers(r, *n, Album, Musica, AlbumConterMusica)
	path, err := r.GetRoute("conter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{album_id}/conter/{musica_id}", path)
	handler := conter(Album, Musica, AlbumConterMusica)
	r.Handle("/v1/{album_id}/conter/{musica_id}", handler)
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/conter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/conter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		AlbumConterMusica.EXPECT().Conter(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/conter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desconter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Album := album.NewMockUseCase(controller)
	AlbumConterMusica := albumcontermusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumConterMusicaHandlers(r, *n, Album, Musica, AlbumConterMusica)
	path, err := r.GetRoute("desconter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{album_id}/desconter/{musica_id}", path)
	handler := desconter(Album, Musica, AlbumConterMusica)
	r.Handle("/v1/{album_id}/desconter/{musica_id}", handler)
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desconter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desconter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		Album.EXPECT().GetAlbum(o.ID).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		AlbumConterMusica.EXPECT().Desconter(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desconter/%d", ts.URL, o.ID, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
