package artistagravarmusica

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/api/presenter"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/artistagravarmusica"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yohanalexander/deezefy-music/entity"
)

func gravar(artistaService artista.UseCase, musicaService musica.UseCase, artistagravarmusicaService artistagravarmusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		artista := vars["artista_email"]
		musica := vars["musica_id"]

		b, err := artistaService.GetArtista(artista)
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

		err = artistagravarmusicaService.Gravar(b, u)
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

func desgravar(artistaService artista.UseCase, musicaService musica.UseCase, artistagravarmusicaService artistagravarmusica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		artista := vars["artista_email"]
		musica := vars["musica_id"]

		b, err := artistaService.GetArtista(artista)
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

		err = artistagravarmusicaService.Desgravar(b, u)
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

// MakeArtistaGravarMusicaHandlers make url handlers
func MakeArtistaGravarMusicaHandlers(r *mux.Router, n negroni.Negroni, artistaService artista.UseCase, musicaService musica.UseCase, artistagravarmusicaService artistagravarmusica.UseCase) {
	r.Handle("/v1/{artista_email}/gravar/{musica_id}", n.With(
		negroni.Wrap(gravar(artistaService, musicaService, artistagravarmusicaService)),
	)).Methods("GET", "OPTIONS").Name("gravar")

	r.Handle("/v1/{artista_email}/desgravar/{musica_id}", n.With(
		negroni.Wrap(desgravar(artistaService, musicaService, artistagravarmusicaService)),
	)).Methods("GET", "OPTIONS").Name("desgravar")
}
