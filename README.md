# Deezefy Music

O projeto [Deezefy](https://deezefy-music.herokuapp.com) consiste de uma API Rest que simula um serviço de músicas, onde estão
implementadas as funções de CRUD (Create, Update, Retrieve, Delete) das entidades
presentes na camada de domínio da aplicação.

Durante o desenvolvimento foi utilizada a metodologia `TDD` (Test Driven Development),
e como padrão de arquitetura a aplicação utiliza o `Clean Architecture`, dividindo
o software em camadas que se comunicam via interfaces e injeção de dependências,
para reduzir o acoplamento e tornar o código mais sustentável.

![Clean Architecture](https://user-images.githubusercontent.com/39287022/107729138-c4750c00-6cce-11eb-9a5c-7f477f8b37d4.png)

## Modelagem dos dados

Para fazer a modelagem dos dados, e dos relacionamentos, foi utilizada a ferramenta
`MySQL Workbench`. O DER (diagrama entidade relacionamento) gerado, pode ser visualizado
em (https://raw.githubusercontent.com/YohanAlexander/deezefy-music/main/docs/der/Deezefy.png). A partir do diagrama foram gerados scripts de inserção de mocks na base de
dados e migração inicial da estrutura das tabelas.

A aplicação não faz uso de ORM (Object Relational Mapper). Todas as querys sql feitas
ao banco de dados são enviadas como strings explícitas ao SGBD por meio dos drivers
da aplicação na camada de infraestrutura, via o padrão de repositório.

## Entidades

* Usuário
* Artista
* Ouvinte
* Música
* Playlist
* Gênero
* Album
* Evento
* Local
* Perfil

## Endpoints

Os endpoints das operações CRUD das entidades seguem o seguinte padrão:

| Método 	| CRUD     	| Endpoint            	|
|--------	|----------	|---------------------	|
| GET    	| List     	| /v1/{entidade}      	|
| POST   	| Create   	| /v1/{entidade}      	|
| GET    	| Retrieve 	| /v1/{entidade}/{id} 	|
| DELETE 	| Delete   	| /v1/{entidade}/{id} 	|

Já os endpoints dos relacionamentos N:N modelados seguem o seguinte padrão:

| Método 	| Endpoint                                 	|
|--------	|------------------------------------------	|
| GET    	| /v1/{id_entidade}/relation/{id_entidade} 	|

## Executando os testes automatizados

```
make test
```

## Executando o serviço localmente

Foi criado um arquivo `docker-compose` para facilitar os testes locais:

```
docker-compose -f build/docker-compose.yaml up
```

## TODO

* [ ] Gerar documentação OpenAPI [Swagger](https://swagger.io) das rotas e métodos usando o [`swag`](https://github.com/swaggo/swag).
* [ ] Integrar serviço de monitoramento e métricas [Prometheus](https://prometheus.io).
* [ ] Adicionar ferramenta de lint do código no Makefile.
* [ ] Fazer upload da imagem docker do servidor para o container registry.
* [ ] Criar mecanismo de CI/CD via Github Actions ou Travis para atualização da imagem mais recente no Heroku.
