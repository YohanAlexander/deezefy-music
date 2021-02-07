package album

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/album/mock"
)

func Test_listAlbums(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumHandlers(r, *n, m)
	path, err := r.GetRoute("listAlbums").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/album", path)
	u := &entity.Album{
		ID: 1,
	}
	m.EXPECT().
		ListAlbums().
		Return([]*entity.Album{u}, nil)
	ts := httptest.NewServer(listAlbums(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listAlbums_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listAlbums(m))
	defer ts.Close()
	m.EXPECT().
		SearchAlbums("123").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?id=123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listAlbums_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Album{
		ID: 1,
	}
	m.EXPECT().
		SearchAlbums("1").
		Return([]*entity.Album{u}, nil)
	ts := httptest.NewServer(listAlbums(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?id=1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createAlbum(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumHandlers(r, *n, m)
	path, err := r.GetRoute("createAlbum").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/album", path)

	m.EXPECT().
		CreateAlbum(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(1, nil)
	h := createAlbum(m)
	s := &entity.Album{
		ID: 1,
	}
	m.EXPECT().
		GetAlbum(s.ID).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"id": 1
}`)
	resp, _ := http.Post(ts.URL+"/v1/album", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Album
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "1", fmt.Sprintf("%d", u.ID))
}

func Test_getAlbum(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumHandlers(r, *n, m)
	path, err := r.GetRoute("getAlbum").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/album/{id}", path)
	u := &entity.Album{
		ID: 1,
	}
	m.EXPECT().
		GetAlbum(u.ID).
		Return(u, nil)
	handler := getAlbum(m)
	r.Handle("/v1/album/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/album/" + strconv.Itoa(u.ID))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Album
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteAlbum(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAlbumHandlers(r, *n, m)
	path, err := r.GetRoute("deleteAlbum").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/album/{id}", path)
	u := &entity.Album{
		ID: 1,
	}
	m.EXPECT().DeleteAlbum(u.ID).Return(nil)
	handler := deleteAlbum(m)
	req, _ := http.NewRequest("DELETE", "/v1/album/"+strconv.Itoa(u.ID), nil)
	r.Handle("/v1/album/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
