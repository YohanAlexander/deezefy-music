package musica

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
)

func Test_listMusicas(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaHandlers(r, *n, m)
	path, err := r.GetRoute("listMusicas").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/musica", path)
	u := &entity.Musica{
		ID: 1,
	}
	m.EXPECT().
		ListMusicas().
		Return([]*entity.Musica{u}, nil)
	ts := httptest.NewServer(listMusicas(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listMusicas_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listMusicas(m))
	defer ts.Close()
	m.EXPECT().
		SearchMusicas("123").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?id=123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listMusicas_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Musica{
		ID: 1,
	}
	m.EXPECT().
		SearchMusicas("1").
		Return([]*entity.Musica{u}, nil)
	ts := httptest.NewServer(listMusicas(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?id=1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createMusica(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaHandlers(r, *n, m)
	path, err := r.GetRoute("createMusica").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/musica", path)

	m.EXPECT().
		CreateMusica(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(1, nil)
	h := createMusica(m)
	s := &entity.Musica{
		ID: 1,
	}
	m.EXPECT().
		GetMusica(s.ID).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"id": 1
}`)
	resp, _ := http.Post(ts.URL+"/v1/musica", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Musica
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "1", fmt.Sprintf("%d", u.ID))
}

func Test_getMusica(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaHandlers(r, *n, m)
	path, err := r.GetRoute("getMusica").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/musica/{id}", path)
	u := &entity.Musica{
		ID: 1,
	}
	m.EXPECT().
		GetMusica(u.ID).
		Return(u, nil)
	handler := getMusica(m)
	r.Handle("/v1/musica/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/musica/" + strconv.Itoa(u.ID))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Musica
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteMusica(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeMusicaHandlers(r, *n, m)
	path, err := r.GetRoute("deleteMusica").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/musica/{id}", path)
	u := &entity.Musica{
		ID: 1,
	}
	m.EXPECT().DeleteMusica(u.ID).Return(nil)
	handler := deleteMusica(m)
	req, _ := http.NewRequest("DELETE", "/v1/musica/"+strconv.Itoa(u.ID), nil)
	r.Handle("/v1/musica/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
