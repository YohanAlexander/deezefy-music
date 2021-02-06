package playlist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/playlist/mock"
)

func Test_listPlaylists(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistHandlers(r, *n, m)
	path, err := r.GetRoute("listPlaylists").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/playlist", path)
	u := &entity.Playlist{
		Nome: "Playlist",
	}
	m.EXPECT().
		ListPlaylists().
		Return([]*entity.Playlist{u}, nil)
	ts := httptest.NewServer(listPlaylists(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listPlaylists_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listPlaylists(m))
	defer ts.Close()
	m.EXPECT().
		SearchPlaylists("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?nome=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listPlaylists_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Playlist{
		Nome: "Playlist",
	}
	m.EXPECT().
		SearchPlaylists("dio").
		Return([]*entity.Playlist{u}, nil)
	ts := httptest.NewServer(listPlaylists(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?nome=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createPlaylist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistHandlers(r, *n, m)
	path, err := r.GetRoute("createPlaylist").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/playlist", path)

	m.EXPECT().
		CreatePlaylist(gomock.Any(), gomock.Any(), gomock.Any()).
		Return("Playlist", nil)
	h := createPlaylist(m)
	s := &entity.Playlist{
		Nome: "Playlist",
	}
	m.EXPECT().
		GetPlaylist(s.Nome).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"nome": "Playlist"
}`)
	resp, _ := http.Post(ts.URL+"/v1/playlist", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Playlist
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "Playlist", fmt.Sprintf("%s", u.Nome))
}

func Test_getPlaylist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistHandlers(r, *n, m)
	path, err := r.GetRoute("getPlaylist").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/playlist/{nome}", path)
	u := &entity.Playlist{
		Nome: "Playlist",
	}
	m.EXPECT().
		GetPlaylist(u.Nome).
		Return(u, nil)
	handler := getPlaylist(m)
	r.Handle("/v1/playlist/{nome}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/playlist/" + u.Nome)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Playlist
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.Nome, d.Nome)
}

func Test_deletePlaylist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePlaylistHandlers(r, *n, m)
	path, err := r.GetRoute("deletePlaylist").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/playlist/{nome}", path)
	u := &entity.Playlist{
		Nome: "Playlist",
	}
	m.EXPECT().DeletePlaylist(u.Nome).Return(nil)
	handler := deletePlaylist(m)
	req, _ := http.NewRequest("DELETE", "/v1/playlist/"+u.Nome, nil)
	r.Handle("/v1/playlist/{nome}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
