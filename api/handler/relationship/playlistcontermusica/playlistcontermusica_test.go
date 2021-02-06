package playlistcontermusica

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
	playlist "github.com/yohanalexander/deezefy-music/usecase/entity/playlist/mock"
	playlistcontermusica "github.com/yohanalexander/deezefy-music/usecase/relationship/playlistcontermusica/mock"
)

func Test_conter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Playlist := playlist.NewMockUseCase(controller)
	PlaylistConterMusica := playlistcontermusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistConterMusicaHandlers(r, *n, Playlist, Musica, PlaylistConterMusica)
	path, err := r.GetRoute("conter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{playlist_nome}/conter/{musica_id}", path)
	handler := conter(Playlist, Musica, PlaylistConterMusica)
	r.Handle("/v1/{playlist_nome}/conter/{musica_id}", handler)
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/conter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/conter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		PlaylistConterMusica.EXPECT().Conter(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/conter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desconter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Playlist := playlist.NewMockUseCase(controller)
	PlaylistConterMusica := playlistcontermusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistConterMusicaHandlers(r, *n, Playlist, Musica, PlaylistConterMusica)
	path, err := r.GetRoute("desconter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{playlist_nome}/desconter/{musica_id}", path)
	handler := desconter(Playlist, Musica, PlaylistConterMusica)
	r.Handle("/v1/{playlist_nome}/desconter/{musica_id}", handler)
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desconter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desconter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		Playlist.EXPECT().GetPlaylist(o.Nome).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		PlaylistConterMusica.EXPECT().Desconter(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desconter/%d", ts.URL, o.Nome, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
