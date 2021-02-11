package usuario

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/usecase/entity/usuario"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listUsuarios(service usuario.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Usuario
		var err error
		email := r.URL.Query().Get("email")

		switch {
		case email == "":
			data, err = service.ListUsuarios()
		default:
			data, err = service.SearchUsuarios(email)
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

		var toJ []presenter.Usuario
		for _, d := range data {
			toJ = presenter.AppendUsuario(*d, toJ)
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

func createUsuario(service usuario.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.PostUsuario{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		email, err := service.CreateUsuario(input.Email, input.Password, input.Birthday)
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

		u, err := service.GetUsuario(email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Usuario{}
		toJ.MakeUsuario(*u)

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

func getUsuario(service usuario.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		email := vars["email"]

		data, err := service.GetUsuario(email)
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

		toJ := &presenter.Usuario{}
		toJ.MakeUsuario(*data)

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

func deleteUsuario(service usuario.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		email := vars["email"]

		err := service.DeleteUsuario(email)

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

// MakeUsuarioHandlers make url handlers
func MakeUsuarioHandlers(r *mux.Router, n negroni.Negroni, service usuario.UseCase) {
	r.Handle("/v1/usuario", n.With(
		negroni.Wrap(listUsuarios(service)),
	)).Methods("GET", "OPTIONS").Name("listUsuarios")

	r.Handle("/v1/usuario", n.With(
		negroni.Wrap(createUsuario(service)),
	)).Methods("POST", "OPTIONS").Name("createUsuario")

	r.Handle("/v1/usuario/{email}", n.With(
		negroni.Wrap(getUsuario(service)),
	)).Methods("GET", "OPTIONS").Name("getUsuario")

	r.Handle("/v1/usuario/{email}", n.With(
		negroni.Wrap(deleteUsuario(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteUsuario")
}
