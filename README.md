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
* Porta 50051 - servidor gRPC


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
# Listagem das Orders
query ListOrders {
  listOrders {
    id
    customer
    total
    items {
      id
      product
      price
      quantity
      total
    }
  }
}

# Mutation para criação de uma Order
mutation CreateOrder {
  createOrder(input: {
    customer: "Roberto Neto"
    items: [
      {
        product:"Produto 1"
        price:10.5
        quantity:2
      }
    ]
  }) {
    id
    customer
    total
    items {
      id
      product
      price
      quantity
      total
    }
  }
}
```