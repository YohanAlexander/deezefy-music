package musica

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listMusicas(service musica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Musica
		var err error
		id := r.URL.Query().Get("id")

		switch {
		case id == "":
			data, err = service.ListMusicas()
		default:
			data, err = service.SearchMusicas(id)
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

		var toJ []presenter.Musica
		for _, d := range data {
			toJ = presenter.AppendMusica(*d, toJ)
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(presenter.Sucesso{
			Result:     toJ,
			StatusCode: http.StatusOK,
		})

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

func createMusica(service musica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Musica{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		music, err := service.CreateMusica(input.Nome, input.Duracao, input.ID)
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

		m, err := service.GetMusica(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Musica{}
		toJ.MakeMusica(*m)

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(presenter.Sucesso{
			Result:     toJ,
			StatusCode: http.StatusCreated,
		})

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

func getMusica(service musica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		data, err := service.GetMusica(id)
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

		toJ := &presenter.Musica{}
		toJ.MakeMusica(*data)

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(presenter.Sucesso{
			Result:     toJ,
			StatusCode: http.StatusOK,
		})

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

func deleteMusica(service musica.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := service.DeleteMusica(id)

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(presenter.Sucesso{
			Result:     presenter.SucessDelete,
			StatusCode: http.StatusOK,
		})

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

// MakeMusicaHandlers make url handlers
func MakeMusicaHandlers(r *mux.Router, n negroni.Negroni, service musica.UseCase) {
	r.Handle("/v1/musica", n.With(
		negroni.Wrap(listMusicas(service)),
	)).Methods("GET", "OPTIONS").Name("listMusicas")

	r.Handle("/v1/musica", n.With(
		negroni.Wrap(createMusica(service)),
	)).Methods("POST", "OPTIONS").Name("createMusica")

	r.Handle("/v1/musica/{id}", n.With(
		negroni.Wrap(getMusica(service)),
	)).Methods("GET", "OPTIONS").Name("getMusica")

	r.Handle("/v1/musica/{id}", n.With(
		negroni.Wrap(deleteMusica(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteMusica")
}
