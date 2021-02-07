package perfilfavoritarartista

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
	artista "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	perfil "github.com/yohanalexander/deezefy-music/usecase/entity/perfil/mock"
	perfilfavoritarartista "github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritarartista/mock"
)

func Test_favoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Artista := artista.NewMockUseCase(controller)
	Perfil := perfil.NewMockUseCase(controller)
	PerfilFavoritarArtista := perfilfavoritarartista.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilFavoritarArtistaHandlers(r, *n, Perfil, Artista, PerfilFavoritarArtista)
	path, err := r.GetRoute("favoritar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{perfil_id}/favoritar/{artista_email}", path)
	handler := favoritar(Perfil, Artista, PerfilFavoritarArtista)
	r.Handle("/v1/{perfil_id}/favoritar/{artista_email}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		PerfilFavoritarArtista.EXPECT().Favoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/favoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desfavoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Artista := artista.NewMockUseCase(controller)
	Perfil := perfil.NewMockUseCase(controller)
	PerfilFavoritarArtista := perfilfavoritarartista.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePerfilFavoritarArtistaHandlers(r, *n, Perfil, Artista, PerfilFavoritarArtista)
	path, err := r.GetRoute("desfavoritar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{perfil_id}/desfavoritar/{artista_email}", path)
	handler := desfavoritar(Perfil, Artista, PerfilFavoritarArtista)
	r.Handle("/v1/{perfil_id}/desfavoritar/{artista_email}", handler)
	t.Run("Perfil not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Perfil{
			ID: 1,
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		Perfil.EXPECT().GetPerfil(o.ID).Return(o, nil)
		Artista.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		PerfilFavoritarArtista.EXPECT().Desfavoritar(a, o).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%d/desfavoritar/%s", ts.URL, o.ID, a.Usuario.Email))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
