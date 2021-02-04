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
	@~/go/bin/mockgen -source=usecase/entity/usuario/interface.go -destination=usecase/entity/usuario/mock/usuario.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/artista/interface.go -destination=usecase/entity/artista/mock/artista.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/ouvinte/interface.go -destination=usecase/entity/ouvinte/mock/ouvinte.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/musica/interface.go -destination=usecase/entity/musica/mock/musica.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/album/interface.go -destination=usecase/entity/album/mock/album.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/playlist/interface.go -destination=usecase/entity/playlist/mock/playlist.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/genero/interface.go -destination=usecase/entity/genero/mock/genero.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/evento/interface.go -destination=usecase/entity/evento/mock/evento.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/local/interface.go -destination=usecase/entity/local/mock/local.go -package=mock
	@~/go/bin/mockgen -source=usecase/entity/perfil/interface.go -destination=usecase/entity/perfil/mock/perfil.go -package=mock

test:
	go test -tags testing ./...
