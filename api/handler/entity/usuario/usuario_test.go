package usuario

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/usuario/mock"
)

func Test_listUsuarios(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUsuarioHandlers(r, *n, m)
	path, err := r.GetRoute("listUsuarios").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/usuario", path)
	u := &entity.Usuario{
		Email: "paperkites@spotify.com",
	}
	m.EXPECT().
		ListUsuarios().
		Return([]*entity.Usuario{u}, nil)
	ts := httptest.NewServer(listUsuarios(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listUsuarios_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listUsuarios(m))
	defer ts.Close()
	m.EXPECT().
		SearchUsuarios("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?email=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listUsuarios_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Usuario{
		Email: "paperkites@spotify.com",
	}
	m.EXPECT().
		SearchUsuarios("paperkites@spotify.com").
		Return([]*entity.Usuario{u}, nil)
	ts := httptest.NewServer(listUsuarios(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?email=paperkites@spotify.com")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createUsuario(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUsuarioHandlers(r, *n, m)
	path, err := r.GetRoute("createUsuario").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/usuario", path)

	m.EXPECT().
		CreateUsuario(gomock.Any(), gomock.Any(), gomock.Any()).
		Return("paperkites@spotify.com", nil)
	h := createUsuario(m)
	s := &entity.Usuario{
		Email: "paperkites@spotify.com",
	}
	m.EXPECT().
		GetUsuario(s.Email).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"email": "paperkites@spotify.com",
"password": "neoncrimson",
"birthday":"2020-01-21"
}`)
	resp, _ := http.Post(ts.URL+"/v1/usuario", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Usuario
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "paperkites@spotify.com", fmt.Sprintf("%s", u.Email))
}

func Test_getUsuario(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUsuarioHandlers(r, *n, m)
	path, err := r.GetRoute("getUsuario").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/usuario/{email}", path)
	u := &entity.Usuario{
		Email: "paperkites@spotify.com",
	}
	m.EXPECT().
		GetUsuario(u.Email).
		Return(u, nil)
	handler := getUsuario(m)
	r.Handle("/v1/usuario/{email}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/usuario/" + u.Email)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Usuario
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.Email, d.Email)
}

func Test_deleteUsuario(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUsuarioHandlers(r, *n, m)
	path, err := r.GetRoute("deleteUsuario").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/usuario/{email}", path)
	u := &entity.Usuario{
		Email: "paperkites@spotify.com",
	}
	m.EXPECT().DeleteUsuario(u.Email).Return(nil)
	handler := deleteUsuario(m)
	req, _ := http.NewRequest("DELETE", "/v1/usuario/"+u.Email, nil)
	r.Handle("/v1/usuario/{email}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
