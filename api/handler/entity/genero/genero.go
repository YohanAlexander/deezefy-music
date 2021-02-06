package genero

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listGeneros(service genero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Genero
		var err error
		nome := r.URL.Query().Get("nome")

		switch {
		case nome == "":
			data, err = service.ListGeneros()
		default:
			data, err = service.SearchGeneros(nome)
		}

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		var toJ []*presenter.Genero
		for _, d := range data {
			toJ = presenter.AppendGenero(*d, toJ)
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

func createGenero(service genero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Genero{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		music, err := service.CreateGenero(input.Nome, input.Estilo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		m, err := service.GetGenero(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Genero{}
		toJ.GetGenero(*m)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

func getGenero(service genero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		nome := vars["nome"]

		data, err := service.GetGenero(nome)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		toJ := &presenter.Genero{}
		toJ.GetGenero(*data)

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

func deleteGenero(service genero.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		nome := vars["nome"]

		err := service.DeleteGenero(nome)

		w.WriteHeader(http.StatusOK)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	})
}

//MakeGeneroHandlers make url handlers
func MakeGeneroHandlers(r *mux.Router, n negroni.Negroni, service genero.UseCase) {
	r.Handle("/v1/genero", n.With(
		negroni.Wrap(listGeneros(service)),
	)).Methods("GET", "OPTIONS").Name("listGeneros")

	r.Handle("/v1/genero", n.With(
		negroni.Wrap(createGenero(service)),
	)).Methods("POST", "OPTIONS").Name("createGenero")

	r.Handle("/v1/genero/{nome}", n.With(
		negroni.Wrap(getGenero(service)),
	)).Methods("GET", "OPTIONS").Name("getGenero")

	r.Handle("/v1/genero/{nome}", n.With(
		negroni.Wrap(deleteGenero(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteGenero")
}
