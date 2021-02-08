# Imagem oficial do golang com suporte a go modules
FROM golang:1.14 AS dev

# Acessando o diretório de trabalho
WORKDIR /app

# Cacheando as depêndencias
COPY go.mod go.sum ./

# Instalando as depêndencias
RUN go mod download

# Instalando compile daemon para live-reload do servidor
RUN GO111MODULE=off go get github.com/githubnemo/CompileDaemon

# Copiando o projeto do host para o container
COPY . .

# Habilitando o modo live-reload
CMD CompileDaemon --build="go build -tags dev -o ./bin/api api/main.go" --command=./bin/api

# Build multi-stage para ambientes de produção
FROM development AS prod

# Compilando o binário
RUN go build -tags prod -o ./bin/api api/main.go

# Habilitando o servidor
CMD ./bin/api
