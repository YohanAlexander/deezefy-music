.PHONY: all
all: build
FORCE: ;

SHELL  := env DEEZEFY_ENV=$(DEEZEFY_ENV) $(SHELL)
DEEZEFY_ENV ?= dev

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api build-cmd

build-api:
	go build -tags $(DEEZEFY_ENV) -o ./bin/api api/main.go

ci: dependencies test

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/usuario/interface.go -destination=usecase/usuario/mock/usuario.go -package=mock
	@~/go/bin/mockgen -source=usecase/artista/interface.go -destination=usecase/artista/mock/artista.go -package=mock
	@~/go/bin/mockgen -source=usecase/ouvinte/interface.go -destination=usecase/ouvinte/mock/ouvinte.go -package=mock
	@~/go/bin/mockgen -source=usecase/musica/interface.go -destination=usecase/musica/mock/musica.go -package=mock
	@~/go/bin/mockgen -source=usecase/album/interface.go -destination=usecase/album/mock/album.go -package=mock
	@~/go/bin/mockgen -source=usecase/playlist/interface.go -destination=usecase/playlist/mock/playlist.go -package=mock
	@~/go/bin/mockgen -source=usecase/genero/interface.go -destination=usecase/genero/mock/genero.go -package=mock
	@~/go/bin/mockgen -source=usecase/evento/interface.go -destination=usecase/evento/mock/evento.go -package=mock
	@~/go/bin/mockgen -source=usecase/local/interface.go -destination=usecase/local/mock/local.go -package=mock
	@~/go/bin/mockgen -source=usecase/perfil/interface.go -destination=usecase/perfil/mock/perfil.go -package=mock
	@~/go/bin/mockgen -source=usecase/segue/interface.go -destination=usecase/segue/mock/segue.go -package=mock
	@~/go/bin/mockgen -source=usecase/curte/interface.go -destination=usecase/curte/mock/curte.go -package=mock
	@~/go/bin/mockgen -source=usecase/salvaplaylist/interface.go -destination=usecase/salvaplaylist/mock/salvaplaylist.go -package=mock
	@~/go/bin/mockgen -source=usecase/salvaalbum/interface.go -destination=usecase/salvaalbum/mock/salvaalbum.go -package=mock
	@~/go/bin/mockgen -source=usecase/grava/interface.go -destination=usecase/grava/mock/grava.go -package=mock
	@~/go/bin/mockgen -source=usecase/artistasfavoritos/interface.go -destination=usecase/artistasfavoritos/mock/artistasfavoritos.go -package=mock
	@~/go/bin/mockgen -source=usecase/artistagenero/interface.go -destination=usecase/artistagenero/mock/artistagenero.go -package=mock
	@~/go/bin/mockgen -source=usecase/ocorre/interface.go -destination=usecase/ocorre/mock/ocorre.go -package=mock
	@~/go/bin/mockgen -source=usecase/musicaplaylist/interface.go -destination=usecase/musicaplaylist/mock/musicaplaylist.go -package=mock
	@~/go/bin/mockgen -source=usecase/albummusica/interface.go -destination=usecase/albummusica/mock/albummusica.go -package=mock
	@~/go/bin/mockgen -source=usecase/musicagenero/interface.go -destination=usecase/musicagenero/mock/musicagenero.go -package=mock
	@~/go/bin/mockgen -source=usecase/generosfavoritos/interface.go -destination=usecase/generosfavoritos/mock/generosfavoritos.go -package=mock
	@~/go/bin/mockgen -source=usecase/criaplaylist/interface.go -destination=usecase/criaplaylist/mock/criaplaylist.go -package=mock

test:
	go test -tags testing ./...
