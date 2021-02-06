package perfilfavoritargenero

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	genero "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
	perfil "github.com/yohanalexander/deezefy-music/usecase/entity/perfil/mock"
	perfilfavoritargenero "github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritargenero/mock"
)

func Test_favoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Perfil := perfil.NewMockUseCase(controller)
	PerfilFavoritarGenero := perfilfavoritargenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilFavoritarGeneroHandlers(r, *n, Perfil, Genero, PerfilFavoritarGenero)
	path, err := r.GetRoute("favoritar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{perfil_id}/favoritar/{genero_nome}", path)
	handler := favoritar(Perfil, Genero, PerfilFavoritarGenero)
	r.Handle("/v1/{perfil_id}/favoritar/{genero_nome}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(a, nil)
		PerfilFavoritarGenero.EXPECT().Favoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desfavoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Genero := genero.NewMockUseCase(controller)
	Perfil := perfil.NewMockUseCase(controller)
	PerfilFavoritarGenero := perfilfavoritargenero.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilFavoritarGeneroHandlers(r, *n, Perfil, Genero, PerfilFavoritarGenero)
	path, err := r.GetRoute("desfavoritar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{perfil_id}/desfavoritar/{genero_nome}", path)
	handler := desfavoritar(Perfil, Genero, PerfilFavoritarGenero)
	r.Handle("/v1/{perfil_id}/desfavoritar/{genero_nome}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(a, nil)
		PerfilFavoritarGenero.EXPECT().Desfavoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
