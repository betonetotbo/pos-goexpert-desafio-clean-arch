# Desafio - Clean Architecture

## Iniciando o banco de dados

Primeiro é necessário iniciar a composição docker com a base de dados:

```bash
cd deployments
docker compose up -d
```

Isso fará com que um servidor MySQL inicie na porta 3306 localmente.

## Criando a base de dados

Na primeira vez que você executar a aplicação, é necessário executar antes a migração.

```bash
go run . migrate
```

Este comando irá conectar-se ao MySQL e criar as tabelas da aplicação.


## Executando a aplicação

Para executar a aplicação:
```bash
go run . start
```

Isso fará com que a aplicação inicie, então 2 servidores irão estar disponíveis:
* Porta 8080 - servidor HTTP
  * No path `/rest`, estão os endpoints RESTfull
  * No path `/graph`, estão os endpoints relacionados ao GraphQL
* Porta 50000 - servidor gRPC