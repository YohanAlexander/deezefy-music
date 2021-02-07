package local

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/usecase/entity/local"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listLocals(service local.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Local
		var err error
		id := r.URL.Query().Get("id")

		switch {
		case id == "":
			data, err = service.ListLocals()
		default:
			data, err = service.SearchLocals(id)
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

		var toJ []*presenter.Local
		for _, d := range data {
			toJ = presenter.AppendLocal(*d, toJ)
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

func createLocal(service local.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Local{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		music, err := service.CreateLocal(input.Cidade, input.Pais, input.ID)
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

		m, err := service.GetLocal(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Local{}
		toJ.GetLocal(*m)

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

func getLocal(service local.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		data, err := service.GetLocal(id)
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

		toJ := &presenter.Local{}
		toJ.GetLocal(*data)

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

func deleteLocal(service local.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := service.DeleteLocal(id)

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

// MakeLocalHandlers make url handlers
func MakeLocalHandlers(r *mux.Router, n negroni.Negroni, service local.UseCase) {
	r.Handle("/v1/local", n.With(
		negroni.Wrap(listLocals(service)),
	)).Methods("GET", "OPTIONS").Name("listLocals")

	r.Handle("/v1/local", n.With(
		negroni.Wrap(createLocal(service)),
	)).Methods("POST", "OPTIONS").Name("createLocal")

	r.Handle("/v1/local/{id}", n.With(
		negroni.Wrap(getLocal(service)),
	)).Methods("GET", "OPTIONS").Name("getLocal")

	r.Handle("/v1/local/{id}", n.With(
		negroni.Wrap(deleteLocal(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteLocal")
}
