package login

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/usecase/entity/usuario"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func loginUsuario(service usuario.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Usuario{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		user, err := service.GetUsuario(input.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrNotFound.Error(),
				StatusCode: http.StatusNotFound,
			})
			return
		}

		err = user.ValidatePassword(input.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnauthorized.Error(),
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		toJ := &presenter.Usuario{}
		toJ.MakeUsuario(*user)

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

// MakeLoginHandlers make url handlers
func MakeLoginHandlers(r *mux.Router, n negroni.Negroni, service usuario.UseCase) {
	r.Handle("/v1/login", n.With(
		negroni.Wrap(loginUsuario(service)),
	)).Methods("POST", "OPTIONS").Name("loginUsuario")
}
