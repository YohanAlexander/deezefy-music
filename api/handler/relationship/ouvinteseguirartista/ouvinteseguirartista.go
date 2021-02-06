package ouvinteseguirartista

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvinteseguirartista"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func seguir(ouvinteService ouvinte.UseCase, artistaService artista.UseCase, ouvinteseguirartistaService ouvinteseguirartista.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		artista := vars["artista_email"]

		b, err := ouvinteService.GetOuvinte(ouvinte)
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

		err = ouvinteseguirartistaService.Seguir(b, u)
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

func desseguir(ouvinteService ouvinte.UseCase, artistaService artista.UseCase, ouvinteseguirartistaService ouvinteseguirartista.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		artista := vars["artista_email"]

		b, err := ouvinteService.GetOuvinte(ouvinte)
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

		err = ouvinteseguirartistaService.Desseguir(b, u)
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

// MakeOuvinteSeguirArtistaHandlers make url handlers
func MakeOuvinteSeguirArtistaHandlers(r *mux.Router, n negroni.Negroni, ouvinteService ouvinte.UseCase, artistaService artista.UseCase, ouvinteseguirartistaService ouvinteseguirartista.UseCase) {
	r.Handle("/v1/{ouvinte_email}/seguir/{artista_email}", n.With(
		negroni.Wrap(seguir(ouvinteService, artistaService, ouvinteseguirartistaService)),
	)).Methods("GET", "OPTIONS").Name("seguir")

	r.Handle("/v1/{ouvinte_email}/desseguir/{artista_email}", n.With(
		negroni.Wrap(desseguir(ouvinteService, artistaService, ouvinteseguirartistaService)),
	)).Methods("GET", "OPTIONS").Name("desseguir")
}
