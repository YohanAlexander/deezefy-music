package ouvinte

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listOuvintes(service ouvinte.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Ouvinte
		var err error
		email := r.URL.Query().Get("email")

		switch {
		case email == "":
			data, err = service.ListOuvintes()
		default:
			data, err = service.SearchOuvintes(email)
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

		var toJ []*presenter.Ouvinte
		for _, d := range data {
			toJ = presenter.AppendOuvinte(*d, toJ)
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	})
}

func createOuvinte(service ouvinte.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Ouvinte{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		email, err := service.CreateOuvinte(input.Usuario.Email, input.Usuario.Password, input.Usuario.Birthday, input.PrimeiroNome, input.Sobrenome)
		if err != nil && err != entity.ErrInvalidEntity {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		if err == entity.ErrInvalidEntity {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrInvalidEntity.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		o, err := service.GetOuvinte(email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Ouvinte{}
		toJ.GetOuvinte(*o)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	})
}

func getOuvinte(service ouvinte.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		email := vars["email"]

		data, err := service.GetOuvinte(email)
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

		toJ := &presenter.Ouvinte{}
		toJ.GetOuvinte(*data)

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
	})
}

func deleteOuvinte(service ouvinte.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		email := vars["email"]

		err := service.DeleteOuvinte(email)

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

// MakeOuvinteHandlers make url handlers
func MakeOuvinteHandlers(r *mux.Router, n negroni.Negroni, service ouvinte.UseCase) {
	r.Handle("/v1/ouvinte", n.With(
		negroni.Wrap(listOuvintes(service)),
	)).Methods("GET", "OPTIONS").Name("listOuvintes")

	r.Handle("/v1/ouvinte", n.With(
		negroni.Wrap(createOuvinte(service)),
	)).Methods("POST", "OPTIONS").Name("createOuvinte")

	r.Handle("/v1/ouvinte/{email}", n.With(
		negroni.Wrap(getOuvinte(service)),
	)).Methods("GET", "OPTIONS").Name("getOuvinte")

	r.Handle("/v1/ouvinte/{email}", n.With(
		negroni.Wrap(deleteOuvinte(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteOuvinte")
}
