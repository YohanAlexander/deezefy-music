package ouvintesalvarplaylist

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
	"github.com/yohanalexander/deezefy-music/usecase/entity/playlist"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvarplaylist"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func salvar(ouvinteService ouvinte.UseCase, playlistService playlist.UseCase, ouvintesalvarplaylistService ouvintesalvarplaylist.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		playlist := vars["playlist_nome"]

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

		u, err := playlistService.GetPlaylist(playlist)
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

		err = ouvintesalvarplaylistService.Salvar(b, u)
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

func dessalvar(ouvinteService ouvinte.UseCase, playlistService playlist.UseCase, ouvintesalvarplaylistService ouvintesalvarplaylist.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		playlist := vars["playlist_nome"]

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

		u, err := playlistService.GetPlaylist(playlist)
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

		err = ouvintesalvarplaylistService.Dessalvar(b, u)
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

// MakeOuvinteSalvarPlaylistHandlers make url handlers
func MakeOuvinteSalvarPlaylistHandlers(r *mux.Router, n negroni.Negroni, ouvinteService ouvinte.UseCase, playlistService playlist.UseCase, ouvintesalvarplaylistService ouvintesalvarplaylist.UseCase) {
	r.Handle("/v1/{ouvinte_email}/salvar/{playlist_nome}", n.With(
		negroni.Wrap(salvar(ouvinteService, playlistService, ouvintesalvarplaylistService)),
	)).Methods("GET", "OPTIONS").Name("salvar")

	r.Handle("/v1/{ouvinte_email}/dessalvar/{playlist_nome}", n.With(
		negroni.Wrap(dessalvar(ouvinteService, playlistService, ouvintesalvarplaylistService)),
	)).Methods("GET", "OPTIONS").Name("dessalvar")
}
