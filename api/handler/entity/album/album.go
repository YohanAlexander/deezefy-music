package album

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yohanalexander/deezefy-music/usecase/entity/album"

	"github.com/yohanalexander/deezefy-music/api/presenter"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listAlbums(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data []*entity.Album
		var err error
		id := r.URL.Query().Get("id")

		switch {
		case id == "":
			data, err = service.ListAlbums()
		default:
			data, err = service.SearchAlbums(id)
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

		var toJ []*presenter.Album
		for _, d := range data {
			toJ = presenter.AppendAlbum(*d, toJ)
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

func createAlbum(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		input := &presenter.Album{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrJSON.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		music, err := service.CreateAlbum(input.Artista.Usuario.Email, input.Artista.Usuario.Password, input.Artista.Usuario.Birthday, input.Artista.NomeArtistico, input.Artista.Biografia, input.Titulo, input.Artista.AnoFormacao, input.AnoLancamento, input.ID)
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

		m, err := service.GetAlbum(music)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(presenter.Erro{
				Message:    presenter.ErrUnexpected.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		toJ := &presenter.Album{}
		toJ.GetAlbum(*m)

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

func getAlbum(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		data, err := service.GetAlbum(id)
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

		toJ := &presenter.Album{}
		toJ.GetAlbum(*data)

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

func deleteAlbum(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := service.DeleteAlbum(id)

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

// MakeAlbumHandlers make url handlers
func MakeAlbumHandlers(r *mux.Router, n negroni.Negroni, service album.UseCase) {
	r.Handle("/v1/album", n.With(
		negroni.Wrap(listAlbums(service)),
	)).Methods("GET", "OPTIONS").Name("listAlbums")

	r.Handle("/v1/album", n.With(
		negroni.Wrap(createAlbum(service)),
	)).Methods("POST", "OPTIONS").Name("createAlbum")

	r.Handle("/v1/album/{id}", n.With(
		negroni.Wrap(getAlbum(service)),
	)).Methods("GET", "OPTIONS").Name("getAlbum")

	r.Handle("/v1/album/{id}", n.With(
		negroni.Wrap(deleteAlbum(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteAlbum")
}
