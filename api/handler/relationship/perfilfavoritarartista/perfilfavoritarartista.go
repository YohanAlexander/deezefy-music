package perfilfavoritarartista

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritarartista"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func favoritar(perfilService perfil.UseCase, artistaService artista.UseCase, perfilfavoritarartistaService perfilfavoritarartista.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		perfil := vars["perfil_id"]
		artista := vars["artista_email"]

		id, err := strconv.Atoi(perfil)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		b, err := perfilService.GetPerfil(id)
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

		u, err := artistaService.GetArtista(artista)
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

		err = perfilfavoritarartistaService.Favoritar(u, b)
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

func desfavoritar(perfilService perfil.UseCase, artistaService artista.UseCase, perfilfavoritarartistaService perfilfavoritarartista.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		perfil := vars["perfil_id"]
		artista := vars["artista_email"]

		id, err := strconv.Atoi(perfil)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		b, err := perfilService.GetPerfil(id)
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

		u, err := artistaService.GetArtista(artista)
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

		err = perfilfavoritarartistaService.Desfavoritar(u, b)
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

// MakePerfilFavoritarArtistaHandlers make url handlers
func MakePerfilFavoritarArtistaHandlers(r *mux.Router, n negroni.Negroni, perfilService perfil.UseCase, artistaService artista.UseCase, perfilfavoritarartistaService perfilfavoritarartista.UseCase) {
	r.Handle("/v1/{perfil_id}/favoritar/{artista_email}", n.With(
		negroni.Wrap(favoritar(perfilService, artistaService, perfilfavoritarartistaService)),
	)).Methods("GET", "OPTIONS").Name("favoritar")

	r.Handle("/v1/{perfil_id}/desfavoritar/{artista_email}", n.With(
		negroni.Wrap(desfavoritar(perfilService, artistaService, perfilfavoritarartistaService)),
	)).Methods("GET", "OPTIONS").Name("desfavoritar")
}
