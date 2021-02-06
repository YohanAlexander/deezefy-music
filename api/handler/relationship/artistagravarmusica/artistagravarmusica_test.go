package artistagravarmusica

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
	musica "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	artistagravarmusica "github.com/yohanalexander/deezefy-music/usecase/relationship/artistagravarmusica/mock"
)

func Test_gravar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Artista := artista.NewMockUseCase(controller)
	ArtistaGravarMusica := artistagravarmusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaGravarMusicaHandlers(r, *n, Artista, Musica, ArtistaGravarMusica)
	path, err := r.GetRoute("gravar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{artista_email}/gravar/{musica_id}", path)
	handler := gravar(Artista, Musica, ArtistaGravarMusica)
	r.Handle("/v1/{artista_email}/gravar/{musica_id}", handler)
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/gravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/gravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		ArtistaGravarMusica.EXPECT().Gravar(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/gravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

func Test_desgravar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	Musica := musica.NewMockUseCase(controller)
	Artista := artista.NewMockUseCase(controller)
	ArtistaGravarMusica := artistagravarmusica.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeArtistaGravarMusicaHandlers(r, *n, Artista, Musica, ArtistaGravarMusica)
	path, err := r.GetRoute("desgravar").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/{artista_email}/desgravar/{musica_id}", path)
	handler := desgravar(Artista, Musica, ArtistaGravarMusica)
	r.Handle("/v1/{artista_email}/desgravar/{musica_id}", handler)
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desgravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desgravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
	t.Run("success", func(t *testing.T) {
		o := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		Artista.EXPECT().GetArtista(o.Usuario.Email).Return(o, nil)
		Musica.EXPECT().GetMusica(m.ID).Return(m, nil)
		ArtistaGravarMusica.EXPECT().Desgravar(o, m).Return(nil)
		ts := httptest.NewServer(r)
		defer ts.Close()
		res, err := http.Get(fmt.Sprintf("%s/v1/%s/desgravar/%d", ts.URL, o.Usuario.Email, m.ID))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
