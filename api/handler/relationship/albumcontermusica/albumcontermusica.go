package albumcontermusica

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/album"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/albumcontermusica"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func conter(albumService album.UseCase, musicaService musica.UseCase, albumcontermusicaService albumcontermusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		album := vars["album_id"]
		musica := vars["musica_id"]

		id, err := strconv.Atoi(album)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		b, err := albumService.GetAlbum(id)
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

		id, err = strconv.Atoi(musica)
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

		err = albumcontermusicaService.Conter(b, u)
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

func desconter(albumService album.UseCase, musicaService musica.UseCase, albumcontermusicaService albumcontermusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		album := vars["album_id"]
		musica := vars["musica_id"]

		id, err := strconv.Atoi(album)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		b, err := albumService.GetAlbum(id)
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

		id, err = strconv.Atoi(musica)
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

		err = albumcontermusicaService.Desconter(b, u)
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

// MakeAlbumConterMusicaHandlers make url handlers
func MakeAlbumConterMusicaHandlers(r *mux.Router, n negroni.Negroni, albumService album.UseCase, musicaService musica.UseCase, albumcontermusicaService albumcontermusica.UseCase) {
	r.Handle("/v1/{album_id}/conter/{musica_id}", n.With(
		negroni.Wrap(conter(albumService, musicaService, albumcontermusicaService)),
	)).Methods("GET", "OPTIONS").Name("conter")

	r.Handle("/v1/{album_id}/desconter/{musica_id}", n.With(
		negroni.Wrap(desconter(albumService, musicaService, albumcontermusicaService)),
	)).Methods("GET", "OPTIONS").Name("desconter")
}
