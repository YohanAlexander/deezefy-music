package perfil

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil/mock"
)

func Test_listPerfils(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilHandlers(r, *n, m)
	path, err := r.GetRoute("listPerfils").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/perfil", path)
	u := &entity.Perfil{
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		},
	}
	m.EXPECT().
		ListPerfils().
		Return([]*entity.Perfil{u}, nil)
	ts := httptest.NewServer(listPerfils(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listPerfils_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listPerfils(m))
	defer ts.Close()
	m.EXPECT().
		SearchPerfils("ouvinte@email.com").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?email=ouvinte@email.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listPerfils_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Perfil{
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		},
	}
	m.EXPECT().
		SearchPerfils("ouvinte@email.com").
		Return([]*entity.Perfil{u}, nil)
	ts := httptest.NewServer(listPerfils(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?email=ouvinte@email.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createPerfil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilHandlers(r, *n, m)
	path, err := r.GetRoute("createPerfil").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/perfil", path)

	m.EXPECT().
		CreatePerfil(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return("ouvinte@email.com", nil)
	h := createPerfil(m)
	s := &entity.Perfil{
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		},
	}
	m.EXPECT().
		GetPerfil(s.Ouvinte.Usuario.Email).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"email": "ouvinte@email.com"
}`)
	resp, _ := http.Post(ts.URL+"/v1/perfil", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Perfil
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "ouvinte@email.com", fmt.Sprintf("%s", u.Ouvinte.Usuario.Email))
}

func Test_getPerfil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilHandlers(r, *n, m)
	path, err := r.GetRoute("getPerfil").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/perfil/{email}", path)
	u := &entity.Perfil{
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		},
	}
	m.EXPECT().
		GetPerfil(u.Ouvinte.Usuario.Email).
		Return(u, nil)
	handler := getPerfil(m)
	r.Handle("/v1/perfil/{email}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/perfil/" + u.Ouvinte.Usuario.Email)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Perfil
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deletePerfil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilHandlers(r, *n, m)
	path, err := r.GetRoute("deletePerfil").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/perfil/{email}", path)
	u := &entity.Perfil{
		Ouvinte: entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		},
	}
	m.EXPECT().DeletePerfil(u.Ouvinte.Usuario.Email).Return(nil)
	handler := deletePerfil(m)
	req, _ := http.NewRequest("DELETE", "/v1/perfil/"+u.Ouvinte.Usuario.Email, nil)
	r.Handle("/v1/perfil/{email}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
