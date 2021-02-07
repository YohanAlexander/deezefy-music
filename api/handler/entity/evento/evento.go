package evento

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/usecase/entity/evento"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listEventos(service evento.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Evento
		var err error
		id := r.URL.Query().Get("id")

		switch {
		case id == "":
			data, err = service.ListEventos()
		default:
			data, err = service.SearchEventos(id)
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

		var toJ []*presenter.Evento
		for _, d := range data {
			toJ = presenter.AppendEvento(*d, toJ)
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

func createEvento(service evento.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Evento{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		music, err := service.CreateEvento(input.Usuario.Email, input.Usuario.Password, input.Usuario.Birthday, input.Nome, input.Data, input.Local.Cidade, input.Local.Pais, input.Local.ID, input.ID)
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

		m, err := service.GetEvento(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Evento{}
		toJ.GetEvento(*m)

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

func getEvento(service evento.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		data, err := service.GetEvento(id)
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

		toJ := &presenter.Evento{}
		toJ.GetEvento(*data)

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

func deleteEvento(service evento.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := service.DeleteEvento(id)

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

// MakeEventoHandlers make url handlers
func MakeEventoHandlers(r *mux.Router, n negroni.Negroni, service evento.UseCase) {
	r.Handle("/v1/evento", n.With(
		negroni.Wrap(listEventos(service)),
	)).Methods("GET", "OPTIONS").Name("listEventos")

	r.Handle("/v1/evento", n.With(
		negroni.Wrap(createEvento(service)),
	)).Methods("POST", "OPTIONS").Name("createEvento")

	r.Handle("/v1/evento/{id}", n.With(
		negroni.Wrap(getEvento(service)),
	)).Methods("GET", "OPTIONS").Name("getEvento")

	r.Handle("/v1/evento/{id}", n.With(
		negroni.Wrap(deleteEvento(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteEvento")
}
