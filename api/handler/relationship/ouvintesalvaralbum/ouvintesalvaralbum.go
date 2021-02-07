package ouvintesalvaralbum

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/album"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvaralbum"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func salvar(ouvinteService ouvinte.UseCase, albumService album.UseCase, ouvintesalvaralbumService ouvintesalvaralbum.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		album := vars["album_id"]

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

		id, err := strconv.Atoi(album)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		u, err := albumService.GetAlbum(id)
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

		err = ouvintesalvaralbumService.Salvar(b, u)
		w.WriteHeader(http.StatusCreated)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	})
}

func dessalvar(ouvinteService ouvinte.UseCase, albumService album.UseCase, ouvintesalvaralbumService ouvintesalvaralbum.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		ouvinte := vars["ouvinte_email"]
		album := vars["album_id"]

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

		id, err := strconv.Atoi(album)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		u, err := albumService.GetAlbum(id)
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

		err = ouvintesalvaralbumService.Dessalvar(b, u)
		w.WriteHeader(http.StatusCreated)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	})
}

// MakeOuvinteSalvarAlbumHandlers make url handlers
func MakeOuvinteSalvarAlbumHandlers(r *mux.Router, n negroni.Negroni, ouvinteService ouvinte.UseCase, albumService album.UseCase, ouvintesalvaralbumService ouvintesalvaralbum.UseCase) {
	r.Handle("/v1/{ouvinte_email}/salvar/{album_id}", n.With(
		negroni.Wrap(salvar(ouvinteService, albumService, ouvintesalvaralbumService)),
	)).Methods("GET", "OPTIONS").Name("salvar")

	r.Handle("/v1/{ouvinte_email}/dessalvar/{album_id}", n.With(
		negroni.Wrap(dessalvar(ouvinteService, albumService, ouvintesalvaralbumService)),
	)).Methods("GET", "OPTIONS").Name("dessalvar")
}
