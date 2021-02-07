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
	assert.Equal(t, "/v1/{perfil_email}/favoritar/{genero_nome}", path)
	handler := favoritar(Perfil, Genero, PerfilFavoritarGenero)
	r.Handle("/v1/{perfil_email}/favoritar/{genero_nome}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/favoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/favoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(a, nil)
		PerfilFavoritarGenero.EXPECT().Favoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/favoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
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
	assert.Equal(t, "/v1/{perfil_email}/desfavoritar/{genero_nome}", path)
	handler := desfavoritar(Perfil, Genero, PerfilFavoritarGenero)
	r.Handle("/v1/{perfil_email}/desfavoritar/{genero_nome}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desfavoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Genero not found", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desfavoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a := &entity.Genero{
			Nome: "Genero",
		}
		Perfil.EXPECT().GetPerfil(o.Ouvinte.Usuario.Email).Return(o, nil)
		Genero.EXPECT().GetGenero(a.Nome).Return(a, nil)
		PerfilFavoritarGenero.EXPECT().Desfavoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desfavoritar/%s", ts.URL, o.Ouvinte.Usuario.Email, a.Nome))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
