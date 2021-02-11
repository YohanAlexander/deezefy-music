package perfil

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listPerfils(service perfil.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Perfil
		var err error
		id := r.URL.Query().Get("id")

		switch {
		case id == "":
			data, err = service.ListPerfils()
		default:
			data, err = service.SearchPerfils(id)
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

		var toJ []*presenter.Perfil
		for _, d := range data {
			toJ = presenter.AppendPerfil(*d, toJ)
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

func createPerfil(service perfil.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Perfil{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		music, err := service.CreatePerfil(input.Ouvinte.Usuario.Email, input.Ouvinte.Usuario.Password, input.Ouvinte.Usuario.Birthday, input.Ouvinte.PrimeiroNome, input.Ouvinte.Sobrenome, input.InformacoesRelevantes, input.ID)
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

		m, err := service.GetPerfil(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Perfil{}
		toJ.GetPerfil(*m)

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

func getPerfil(service perfil.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		data, err := service.GetPerfil(id)
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

		toJ := &presenter.Perfil{}
		toJ.GetPerfil(*data)

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

func deletePerfil(service perfil.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := service.DeletePerfil(id)

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

// MakePerfilHandlers make url handlers
func MakePerfilHandlers(r *mux.Router, n negroni.Negroni, service perfil.UseCase) {
	r.Handle("/v1/perfil", n.With(
		negroni.Wrap(listPerfils(service)),
	)).Methods("GET", "OPTIONS").Name("listPerfils")

	r.Handle("/v1/perfil", n.With(
		negroni.Wrap(createPerfil(service)),
	)).Methods("POST", "OPTIONS").Name("createPerfil")

	r.Handle("/v1/perfil/{id}", n.With(
		negroni.Wrap(getPerfil(service)),
	)).Methods("GET", "OPTIONS").Name("getPerfil")

	r.Handle("/v1/perfil/{id}", n.With(
		negroni.Wrap(deletePerfil(service)),
	)).Methods("DELETE", "OPTIONS").Name("deletePerfil")
}
