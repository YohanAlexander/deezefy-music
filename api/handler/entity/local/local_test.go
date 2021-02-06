package local

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
	"github.com/yohanalexander/deezefy-music/usecase/entity/local/mock"
)

func Test_listLocals(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLocalHandlers(r, *n, m)
	path, err := r.GetRoute("listLocals").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/local", path)
	u := &entity.Local{
		ID: 1,
	}
	m.EXPECT().
		ListLocals().
		Return([]*entity.Local{u}, nil)
	ts := httptest.NewServer(listLocals(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listLocals_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listLocals(m))
	defer ts.Close()
	m.EXPECT().
		SearchLocals("123").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?id=123")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listLocals_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Local{
		ID: 1,
	}
	m.EXPECT().
		SearchLocals("1").
		Return([]*entity.Local{u}, nil)
	ts := httptest.NewServer(listLocals(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?id=1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createLocal(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLocalHandlers(r, *n, m)
	path, err := r.GetRoute("createLocal").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/local", path)

	m.EXPECT().
		CreateLocal(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(1, nil)
	h := createLocal(m)
	s := &entity.Local{
		ID: 1,
	}
	m.EXPECT().
		GetLocal(s.ID).
		Return(s, nil)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
"id": 1
}`)
	resp, _ := http.Post(ts.URL+"/v1/local", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Local
	json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "1", fmt.Sprintf("%d", u.ID))
}

func Test_getLocal(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLocalHandlers(r, *n, m)
	path, err := r.GetRoute("getLocal").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/local/{id}", path)
	u := &entity.Local{
		ID: 1,
	}
	m.EXPECT().
		GetLocal(u.ID).
		Return(u, nil)
	handler := getLocal(m)
	r.Handle("/v1/local/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/local/" + strconv.Itoa(u.ID))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Local
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteLocal(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeLocalHandlers(r, *n, m)
	path, err := r.GetRoute("deleteLocal").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/local/{id}", path)
	u := &entity.Local{
		ID: 1,
	}
	m.EXPECT().DeleteLocal(u.ID).Return(nil)
	handler := deleteLocal(m)
	req, _ := http.NewRequest("DELETE", "/v1/local/"+strconv.Itoa(u.ID), nil)
	r.Handle("/v1/local/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
