package ouvinte

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
)

func Test_listOuvintes(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteHandlers(r, *n, m)
	path, err := r.GetRoute("listOuvintes").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/ouvinte", path)
	o := &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		ListOuvintes().
		Return([]*entity.Ouvinte{o}, nil)
	ts := httptest.NewServer(listOuvintes(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listOuvintes_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listOuvintes(m))
	defer ts.Close()
	m.EXPECT().
		SearchOuvintes("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?email=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listOuvintes_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	o := &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		SearchOuvintes("paperkites@spotify.com").
		Return([]*entity.Ouvinte{o}, nil)
	ts := httptest.NewServer(listOuvintes(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?email=paperkites@spotify.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createOuvinte(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteHandlers(r, *n, m)
	path, err := r.GetRoute("createOuvinte").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/ouvinte", path)

	m.EXPECT().
		CreateOuvinte(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return("paperkites@spotify.com", nil)
	h := createOuvinte(m)
	s := &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		GetOuvinte(s.Usuario.Email).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"email": "paperkites@spotify.com",
"password": "neoncrimson",
"birthday":"2020-01-21"
}`)
	resp, _ := http.Post(ts.URL+"/v1/ouvinte", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var o *presenter.Ouvinte
	json.NewDecoder(resp.Body).Decode(&o)
	assert.Equal(t, "paperkites@spotify.com", fmt.Sprintf("%s", o.Usuario.Email))
}

func Test_getOuvinte(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteHandlers(r, *n, m)
	path, err := r.GetRoute("getOuvinte").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/ouvinte/{email}", path)
	o := &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		GetOuvinte(o.Usuario.Email).
		Return(o, nil)
	handler := getOuvinte(m)
	r.Handle("/v1/ouvinte/{email}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/ouvinte/" + o.Usuario.Email)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Ouvinte
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, o.Usuario.Email, d.Usuario.Email)
}

func Test_deleteOuvinte(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeOuvinteHandlers(r, *n, m)
	path, err := r.GetRoute("deleteOuvinte").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/ouvinte/{email}", path)
	o := &entity.Ouvinte{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().DeleteOuvinte(o.Usuario.Email).Return(nil)
	handler := deleteOuvinte(m)
	req, _ := http.NewRequest("DELETE", "/v1/ouvinte/"+o.Usuario.Email, nil)
	r.Handle("/v1/ouvinte/{email}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
