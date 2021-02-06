package evento

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/evento/mock"
)

func Test_listEventos(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeEventoHandlers(r, *n, m)
	path, err := r.GetRoute("listEventos").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/evento", path)
	u := &entity.Evento{
		ID: 1,
	}
	m.EXPECT().
		ListEventos().
		Return([]*entity.Evento{u}, nil)
	ts := httptest.NewServer(listEventos(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listEventos_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listEventos(m))
	defer ts.Close()
	m.EXPECT().
		SearchEventos("123").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?id=123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listEventos_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Evento{
		ID: 1,
	}
	m.EXPECT().
		SearchEventos("1").
		Return([]*entity.Evento{u}, nil)
	ts := httptest.NewServer(listEventos(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?id=1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createEvento(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeEventoHandlers(r, *n, m)
	path, err := r.GetRoute("createEvento").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/evento", path)

	m.EXPECT().
		CreateEvento(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(1, nil)
	h := createEvento(m)
	s := &entity.Evento{
		ID: 1,
	}
	m.EXPECT().
		GetEvento(s.ID).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"id": 1
}`)
	resp, _ := http.Post(ts.URL+"/v1/evento", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Evento
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "1", fmt.Sprintf("%d", u.ID))
}

func Test_getEvento(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeEventoHandlers(r, *n, m)
	path, err := r.GetRoute("getEvento").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/evento/{id}", path)
	u := &entity.Evento{
		ID: 1,
	}
	m.EXPECT().
		GetEvento(u.ID).
		Return(u, nil)
	handler := getEvento(m)
	r.Handle("/v1/evento/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/evento/" + strconv.Itoa(u.ID))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Evento
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteEvento(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeEventoHandlers(r, *n, m)
	path, err := r.GetRoute("deleteEvento").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/evento/{id}", path)
	u := &entity.Evento{
		ID: 1,
	}
	m.EXPECT().DeleteEvento(u.ID).Return(nil)
	handler := deleteEvento(m)
	req, _ := http.NewRequest("DELETE", "/v1/evento/"+strconv.Itoa(u.ID), nil)
	r.Handle("/v1/evento/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
