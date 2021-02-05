package artista

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
)

func Test_listArtistas(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaHandlers(r, *n, m)
	path, err := r.GetRoute("listArtistas").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/artista", path)
	a := &entity.Artista{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		ListArtistas().
		Return([]*entity.Artista{a}, nil)
	ts := httptest.NewServer(listArtistas(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listArtistas_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listArtistas(m))
	defer ts.Close()
	m.EXPECT().
		SearchArtistas("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?email=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listArtistas_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	a := &entity.Artista{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		SearchArtistas("paperkites@spotify.com").
		Return([]*entity.Artista{a}, nil)
	ts := httptest.NewServer(listArtistas(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?email=paperkites@spotify.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createArtista(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaHandlers(r, *n, m)
	path, err := r.GetRoute("createArtista").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/artista", path)

	m.EXPECT().
		CreateArtista(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return("paperkites@spotify.com", nil)
	h := createArtista(m)
	s := &entity.Artista{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		GetArtista(s.Usuario.Email).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"email": "paperkites@spotify.com",
"password": "neoncrimson",
"birthday":"2020-01-21"
}`)
	resp, _ := http.Post(ts.URL+"/v1/artista", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var a *presenter.Artista
	json.NewDecoder(resp.Body).Decode(&a)
	assert.Equal(t, "paperkites@spotify.com", fmt.Sprintf("%s", a.Usuario.Email))
}

func Test_getArtista(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaHandlers(r, *n, m)
	path, err := r.GetRoute("getArtista").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/artista/{email}", path)
	a := &entity.Artista{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().
		GetArtista(a.Usuario.Email).
		Return(a, nil)
	handler := getArtista(m)
	r.Handle("/v1/artista/{email}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/artista/" + a.Usuario.Email)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Artista
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, a.Usuario.Email, d.Usuario.Email)
}

func Test_deleteArtista(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaHandlers(r, *n, m)
	path, err := r.GetRoute("deleteArtista").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/artista/{email}", path)
	a := &entity.Artista{
		Usuario: entity.Usuario{
			Email: "paperkites@spotify.com",
		},
	}
	m.EXPECT().DeleteArtista(a.Usuario.Email).Return(nil)
	handler := deleteArtista(m)
	req, _ := http.NewRequest("DELETE", "/v1/artista/"+a.Usuario.Email, nil)
	r.Handle("/v1/artista/{email}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
