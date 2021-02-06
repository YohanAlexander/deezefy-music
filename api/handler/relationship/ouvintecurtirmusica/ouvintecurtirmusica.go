package ouvintecurtirmusica

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintecurtirmusica"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func curtir(ouvinteService ouvinte.UseCase, musicaService musica.UseCase, ouvintecurtirmusicaService ouvintecurtirmusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		musica := vars["musica_id"]

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

		id, err := strconv.Atoi(musica)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		u, err := musicaService.GetMusica(id)
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

		err = ouvintecurtirmusicaService.Curtir(b, u)
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

func descurtir(ouvinteService ouvinte.UseCase, musicaService musica.UseCase, ouvintecurtirmusicaService ouvintecurtirmusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		musica := vars["musica_id"]

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

		id, err := strconv.Atoi(musica)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		u, err := musicaService.GetMusica(id)
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

		err = ouvintecurtirmusicaService.Descurtir(b, u)
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

// MakeOuvinteCurtirMusicaHandlers make url handlers
func MakeOuvinteCurtirMusicaHandlers(r *mux.Router, n negroni.Negroni, ouvinteService ouvinte.UseCase, musicaService musica.UseCase, ouvintecurtirmusicaService ouvintecurtirmusica.UseCase) {
	r.Handle("/v1/{ouvinte_email}/curtir/{musica_id}", n.With(
		negroni.Wrap(curtir(ouvinteService, musicaService, ouvintecurtirmusicaService)),
	)).Methods("GET", "OPTIONS").Name("curtir")

	r.Handle("/v1/{ouvinte_email}/descurtir/{musica_id}", n.With(
		negroni.Wrap(descurtir(ouvinteService, musicaService, ouvintecurtirmusicaService)),
	)).Methods("GET", "OPTIONS").Name("descurtir")
}
