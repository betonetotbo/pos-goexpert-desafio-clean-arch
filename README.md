# Desafio - Clean Architecture

## Detalhes da Implementação

- O aplicativo executa utilizando o cobra-cli.
  - As migrações são aplicadas pelo aplicativo via argumento `migrate`.
  - Para iniciar a aplicação use o argumento `start`.
- O servidor HTTP é o mesmo compartilhado entre os endpoints REST e GraphQL.
- O servidor gRPC roda em outra porta.
- As dependências são mantidas e compartilhadas por um `Value` em um `context.Context` que é configurado na 
  inicialização do cobra-cli. Instâncias como: 
  - A instância do `sql.DB`
  - A instância do `config.Configuration`
- A camada de persistência (repository) possui uma implementação utilizando o `sqlc`.
- A camada de negócios (entity e usecase) são compartilhados no REST, GraphQL e gRPC.

## Configurações

As configurações da aplicação estão no arquivo `.env` na raiz do projeto.
Todos os comandos a seguir devem ser executados dentro da raiz do projeto.

## Iniciando o banco de dados

Primeiro é necessário iniciar a composição docker com a base de dados:

```bash
cd docker
docker compose up -d
```

Isso fará com que um servidor MySQL inicie na porta 3306 localmente.

## Criando a base de dados

Na primeira vez que você executar a aplicação, é necessário executar as migrações.

```bash
go run main.go migrate
```

Este comando irá conectar-se ao MySQL e criar as tabelas da aplicação.

## Executando a aplicação

Para executar a aplicação:
```bash
go run main.go start
```

Isso fará com que a aplicação inicie, então 2 servidores irão estar disponíveis:
* Porta 8080 - servidor HTTP
  * No path `/rest`, estão os endpoints RESTfull
  * No path `/graph`, estão os endpoints relacionados ao GraphQL
* Porta 50051 - servidor gRPC

> Para demais parâmetros (ex.: informar endereço do mysql ou trocar a porta HTTP) execute `go run . help start`

## Testando HTTP Rest

O arquivo `scripts/api.http` tem exemplos de chamadas para listagem de Orders e criação de Orders.

## Testando  gRPC

Você pode utilizar o client do evans, executando o script `scripts/client-grpc.sh`.

Alguns exemplos de comandos para utilizar dentro do client evans:
- `service OrderService` - comando para setar o serviço de Orders
- `call ListOrders` - comando para listagem das Orders
- `call CreateOrder` - comando para criação de Order

## Testando o GraphQL

Acesse a URL http://localhost:8080/graph (o playground do GraphQL).

```graphql
query ListOrders {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}

mutation CreateOrder {
  createOrder(input: {
    price: 10.4
    tax: 1.99
  }) {
    id
    price
    tax
    finalPrice
  }
}
```