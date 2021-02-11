package login

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
	"github.com/yohanalexander/deezefy-music/usecase/entity/usuario"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func loginUsuario(usuarioService usuario.UseCase, ouvinteService ouvinte.UseCase, artistaService artista.UseCase) http.Handler {
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

		toJ := &presenter.Login{}

		o, err := ouvinteService.GetOuvinte(input.Email)
		toJ.Ouvinte = &presenter.Sucesso{
			Result:     o,
			StatusCode: http.StatusOK,
		}

		erro := loginOuvinte(ouvinteService, input, toJ)

		a, err := artistaService.GetArtista(input.Email)
		toJ.Artista = &presenter.Sucesso{
			Result:     a,
			StatusCode: http.StatusOK,
		}

		erra := loginArtista(artistaService, input, toJ)

		if erro == presenter.ErrUnexpected && erra == presenter.ErrUnexpected {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if erro == presenter.ErrUnexpected && erra == presenter.ErrUnexpected {
			w.WriteHeader(http.StatusNotFound)
		}

		if erro == presenter.ErrUnauthorized && erra == presenter.ErrUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(presenter.Login{
			Artista: toJ.Artista,
			Ouvinte: toJ.Ouvinte,
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

func loginOuvinte(ouvinteService ouvinte.UseCase, input *presenter.Usuario, login *presenter.Login) error {

	ouvinte, err := ouvinteService.GetOuvinte(input.Email)
	if err != nil {
		login.Ouvinte = presenter.Erro{
			Message:    presenter.ErrUnexpected.Error(),
			StatusCode: http.StatusInternalServerError,
		}
		return presenter.ErrUnexpected
	}

	if ouvinte == nil {
		login.Ouvinte = presenter.Erro{
			Message:    presenter.ErrNotFound.Error(),
			StatusCode: http.StatusNotFound,
		}
		return presenter.ErrNotFound
	}

	err = ouvinte.Usuario.ValidatePassword(input.Password)
	if err != nil {
		login.Ouvinte = presenter.Erro{
			Message:    presenter.ErrUnauthorized.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		return presenter.ErrUnauthorized
	}
	return nil
}

func loginArtista(artistaService artista.UseCase, input *presenter.Usuario, login *presenter.Login) error {

	artista, err := artistaService.GetArtista(input.Email)
	if err != nil {
		login.Artista = presenter.Erro{
			Message:    presenter.ErrUnexpected.Error(),
			StatusCode: http.StatusInternalServerError,
		}
		return presenter.ErrUnexpected
	}

	if artista == nil {
		login.Artista = presenter.Erro{
			Message:    presenter.ErrNotFound.Error(),
			StatusCode: http.StatusNotFound,
		}
		return presenter.ErrNotFound
	}

	err = artista.Usuario.ValidatePassword(input.Password)
	if err != nil {
		login.Artista = presenter.Erro{
			Message:    presenter.ErrUnauthorized.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		return presenter.ErrUnauthorized
	}
	return nil
}

// MakeLoginHandlers make url handlers
func MakeLoginHandlers(r *mux.Router, n negroni.Negroni, usuarioService usuario.UseCase, ouvinteService ouvinte.UseCase, artistaService artista.UseCase) {
	r.Handle("/v1/login", n.With(
		negroni.Wrap(loginUsuario(usuarioService, ouvinteService, artistaService)),
	)).Methods("POST", "OPTIONS").Name("loginUsuario")
}
