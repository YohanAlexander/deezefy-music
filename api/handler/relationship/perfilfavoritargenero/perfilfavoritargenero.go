package perfilfavoritargenero

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritargenero"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func favoritar(perfilService perfil.UseCase, generoService genero.UseCase, perfilfavoritargeneroService perfilfavoritargenero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		perfil := vars["perfil_email"]
		genero := vars["genero_nome"]

		b, err := perfilService.GetPerfil(perfil)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		if b == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		u, err := generoService.GetGenero(genero)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		if u == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		err = perfilfavoritargeneroService.Favoritar(u, b)
		w.WriteHeader(http.StatusCreated)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

func desfavoritar(perfilService perfil.UseCase, generoService genero.UseCase, perfilfavoritargeneroService perfilfavoritargenero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		perfil := vars["perfil_email"]
		genero := vars["genero_nome"]

		b, err := perfilService.GetPerfil(perfil)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		if b == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		u, err := generoService.GetGenero(genero)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		if u == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		err = perfilfavoritargeneroService.Desfavoritar(u, b)
		w.WriteHeader(http.StatusCreated)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

// MakePerfilFavoritarGeneroHandlers make url handlers
func MakePerfilFavoritarGeneroHandlers(r *mux.Router, n negroni.Negroni, perfilService perfil.UseCase, generoService genero.UseCase, perfilfavoritargeneroService perfilfavoritargenero.UseCase) {
	r.Handle("/v1/{perfil_email}/favoritar/{genero_nome}", n.With(
		negroni.Wrap(favoritar(perfilService, generoService, perfilfavoritargeneroService)),
	)).Methods("GET", "OPTIONS").Name("favoritar")

	r.Handle("/v1/{perfil_email}/desfavoritar/{genero_nome}", n.With(
		negroni.Wrap(desfavoritar(perfilService, generoService, perfilfavoritargeneroService)),
	)).Methods("GET", "OPTIONS").Name("desfavoritar")
}
