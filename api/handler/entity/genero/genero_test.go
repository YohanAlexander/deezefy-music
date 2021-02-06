package genero

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
)

func Test_listGeneros(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeGeneroHandlers(r, *n, m)
	path, err := r.GetRoute("listGeneros").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/genero", path)
	u := &entity.Genero{
		Nome: "Genero",
	}
	m.EXPECT().
		ListGeneros().
		Return([]*entity.Genero{u}, nil)
	ts := httptest.NewServer(listGeneros(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listGeneros_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listGeneros(m))
	defer ts.Close()
	m.EXPECT().
		SearchGeneros("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?nome=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listGeneros_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Genero{
		Nome: "Genero",
	}
	m.EXPECT().
		SearchGeneros("dio").
		Return([]*entity.Genero{u}, nil)
	ts := httptest.NewServer(listGeneros(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?nome=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createGenero(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeGeneroHandlers(r, *n, m)
	path, err := r.GetRoute("createGenero").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/genero", path)

	m.EXPECT().
		CreateGenero(gomock.Any(), gomock.Any()).
		Return("Genero", nil)
	h := createGenero(m)
	s := &entity.Genero{
		Nome: "Genero",
	}
	m.EXPECT().
		GetGenero(s.Nome).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"nome": "Genero"
}`)
	resp, _ := http.Post(ts.URL+"/v1/genero", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Genero
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "Genero", fmt.Sprintf("%s", u.Nome))
}

func Test_getGenero(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeGeneroHandlers(r, *n, m)
	path, err := r.GetRoute("getGenero").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/genero/{nome}", path)
	u := &entity.Genero{
		Nome: "Genero",
	}
	m.EXPECT().
		GetGenero(u.Nome).
		Return(u, nil)
	handler := getGenero(m)
	r.Handle("/v1/genero/{nome}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/genero/" + u.Nome)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Genero
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.Nome, d.Nome)
}

func Test_deleteGenero(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeGeneroHandlers(r, *n, m)
	path, err := r.GetRoute("deleteGenero").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/genero/{nome}", path)
	u := &entity.Genero{
		Nome: "Genero",
	}
	m.EXPECT().DeleteGenero(u.Nome).Return(nil)
	handler := deleteGenero(m)
	req, _ := http.NewRequest("DELETE", "/v1/genero/"+u.Nome, nil)
	r.Handle("/v1/genero/{nome}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
