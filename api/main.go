package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yohanalexander/deezefy-music/infrastructure/postgres/repository"
	"github.com/yohanalexander/deezefy-music/pkg/exit"

	"github.com/yohanalexander/deezefy-music/usecase/entity/album"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/evento"
	"github.com/yohanalexander/deezefy-music/usecase/entity/genero"
	"github.com/yohanalexander/deezefy-music/usecase/entity/local"
	"github.com/yohanalexander/deezefy-music/usecase/entity/musica"
	"github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte"
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"
	"github.com/yohanalexander/deezefy-music/usecase/entity/playlist"
	"github.com/yohanalexander/deezefy-music/usecase/entity/usuario"

	"github.com/yohanalexander/deezefy-music/usecase/relationship/albumcontermusica"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/artistagravarmusica"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/artistapossuirgenero"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/musicapossuirgenero"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintecurtirmusica"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvaralbum"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvintesalvarplaylist"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/ouvinteseguirartista"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritarartista"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/perfilfavoritargenero"
	"github.com/yohanalexander/deezefy-music/usecase/relationship/playlistcontermusica"

	halbum "github.com/yohanalexander/deezefy-music/api/handler/entity/album"
	hartista "github.com/yohanalexander/deezefy-music/api/handler/entity/artista"
	hevento "github.com/yohanalexander/deezefy-music/api/handler/entity/evento"
	hgenero "github.com/yohanalexander/deezefy-music/api/handler/entity/genero"
	hlocal "github.com/yohanalexander/deezefy-music/api/handler/entity/local"
	hmusica "github.com/yohanalexander/deezefy-music/api/handler/entity/musica"
	houvinte "github.com/yohanalexander/deezefy-music/api/handler/entity/ouvinte"
	hperfil "github.com/yohanalexander/deezefy-music/api/handler/entity/perfil"
	hplaylist "github.com/yohanalexander/deezefy-music/api/handler/entity/playlist"
	husuario "github.com/yohanalexander/deezefy-music/api/handler/entity/usuario"

	halbumcontermusica "github.com/yohanalexander/deezefy-music/api/handler/relationship/albumcontermusica"
	hartistagravarmusica "github.com/yohanalexander/deezefy-music/api/handler/relationship/artistagravarmusica"
	hartistapossuirgenero "github.com/yohanalexander/deezefy-music/api/handler/relationship/artistapossuirgenero"
	hmusicapossuirgenero "github.com/yohanalexander/deezefy-music/api/handler/relationship/musicapossuirgenero"
	houvintecurtirmusica "github.com/yohanalexander/deezefy-music/api/handler/relationship/ouvintecurtirmusica"
	houvintesalvaralbum "github.com/yohanalexander/deezefy-music/api/handler/relationship/ouvintesalvaralbum"
	houvintesalvarplaylist "github.com/yohanalexander/deezefy-music/api/handler/relationship/ouvintesalvarplaylist"
	houvinteseguirartista "github.com/yohanalexander/deezefy-music/api/handler/relationship/ouvinteseguirartista"
	hperfilfavoritarartista "github.com/yohanalexander/deezefy-music/api/handler/relationship/perfilfavoritarartista"
	hperfilfavoritargenero "github.com/yohanalexander/deezefy-music/api/handler/relationship/perfilfavoritargenero"
	hplaylistcontermusica "github.com/yohanalexander/deezefy-music/api/handler/relationship/playlistcontermusica"

	"github.com/yohanalexander/deezefy-music/api/middleware"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yohanalexander/deezefy-music/config"
)

func main() {

	db, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		//negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)

	// usuario
	usuarioRepo := repository.NewUsuarioPSQL(db)
	usuarioService := usuario.NewService(usuarioRepo)
	husuario.MakeUsuarioHandlers(r, *n, usuarioService)

	// artista
	artistaRepo := repository.NewArtistaPSQL(db)
	artistaService := artista.NewService(artistaRepo)
	hartista.MakeArtistaHandlers(r, *n, artistaService)

	// ouvinte
	ouvinteRepo := repository.NewOuvintePSQL(db)
	ouvinteService := ouvinte.NewService(ouvinteRepo)
	houvinte.MakeOuvinteHandlers(r, *n, ouvinteService)

	// musica
	musicaRepo := repository.NewMusicaPSQL(db)
	musicaService := musica.NewService(musicaRepo)
	hmusica.MakeMusicaHandlers(r, *n, musicaService)

	// album
	albumRepo := repository.NewAlbumPSQL(db)
	albumService := album.NewService(albumRepo)
	halbum.MakeAlbumHandlers(r, *n, albumService)

	// playlist
	playlistRepo := repository.NewPlaylistPSQL(db)
	playlistService := playlist.NewService(playlistRepo)
	hplaylist.MakePlaylistHandlers(r, *n, playlistService)

	// evento
	eventoRepo := repository.NewEventoPSQL(db)
	eventoService := evento.NewService(eventoRepo)
	hevento.MakeEventoHandlers(r, *n, eventoService)

	// local
	localRepo := repository.NewLocalPSQL(db)
	localService := local.NewService(localRepo)
	hlocal.MakeLocalHandlers(r, *n, localService)

	// perfil
	perfilRepo := repository.NewPerfilPSQL(db)
	perfilService := perfil.NewService(perfilRepo)
	hperfil.MakePerfilHandlers(r, *n, perfilService)

	// genero
	generoRepo := repository.NewGeneroPSQL(db)
	generoService := genero.NewService(generoRepo)
	hgenero.MakeGeneroHandlers(r, *n, generoService)

	// albumcontermusica
	albumcontermusicaService := albumcontermusica.NewService(albumService, musicaService)
	halbumcontermusica.MakeAlbumConterMusicaHandlers(r, *n, albumService, musicaService, albumcontermusicaService)

	// artistagravarmusica
	artistagravarmusicaService := artistagravarmusica.NewService(artistaService, musicaService)
	hartistagravarmusica.MakeArtistaGravarMusicaHandlers(r, *n, artistaService, musicaService, artistagravarmusicaService)

	// artistapossuirgenero
	artistapossuirgeneroService := artistapossuirgenero.NewService(artistaService, generoService)
	hartistapossuirgenero.MakeArtistaPossuirGeneroHandlers(r, *n, artistaService, generoService, artistapossuirgeneroService)

	// musicapossuirgenero
	musicapossuirgeneroService := musicapossuirgenero.NewService(musicaService, generoService)
	hmusicapossuirgenero.MakeMusicaPossuirGeneroHandlers(r, *n, musicaService, generoService, musicapossuirgeneroService)

	// ouvintecurtirmusica
	ouvintecurtirmusicaService := ouvintecurtirmusica.NewService(ouvinteService, musicaService)
	houvintecurtirmusica.MakeOuvinteCurtirMusicaHandlers(r, *n, ouvinteService, musicaService, ouvintecurtirmusicaService)

	// ouvintesalvaralbum
	ouvintesalvaralbumService := ouvintesalvaralbum.NewService(ouvinteService, albumService)
	houvintesalvaralbum.MakeOuvinteSalvarAlbumHandlers(r, *n, ouvinteService, albumService, ouvintesalvaralbumService)

	// ouvintesalvarplaylist
	ouvintesalvarplaylistService := ouvintesalvarplaylist.NewService(ouvinteService, playlistService)
	houvintesalvarplaylist.MakeOuvinteSalvarPlaylistHandlers(r, *n, ouvinteService, playlistService, ouvintesalvarplaylistService)

	// ouvinteseguirartista
	ouvinteseguirartistaService := ouvinteseguirartista.NewService(ouvinteService, artistaService)
	houvinteseguirartista.MakeOuvinteSeguirArtistaHandlers(r, *n, ouvinteService, artistaService, ouvinteseguirartistaService)

	// perfilfavoritarartista
	perfilfavoritarartistaService := perfilfavoritarartista.NewService(artistaService, perfilService)
	hperfilfavoritarartista.MakePerfilFavoritarArtistaHandlers(r, *n, perfilService, artistaService, perfilfavoritarartistaService)

	// perfilfavoritargenero
	perfilfavoritargeneroService := perfilfavoritargenero.NewService(generoService, perfilService)
	hperfilfavoritargenero.MakePerfilFavoritarGeneroHandlers(r, *n, perfilService, generoService, perfilfavoritargeneroService)

	// playlistcontermusica
	playlistcontermusicaService := playlistcontermusica.NewService(playlistService, musicaService)
	hplaylistcontermusica.MakePlaylistConterMusicaHandlers(r, *n, playlistService, musicaService, playlistcontermusicaService)

	http.Handle("/", r)
	//http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Servidor OK!")
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.API_PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	go func() {
		log.Println("Iniciando servidor na porta:", config.API_PORT)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	exit.Init(func() {
		if err := srv.Close(); err != nil {
			log.Println(err.Error())
		}
		if err := db.Close(); err != nil {
			log.Println(err.Error())
		}
	})

}
