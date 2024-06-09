# desafio-clean-architecture

## Pré-Requisitos

- GOLANG 1.20;
- [Composer](https://getcomposer.org);
- [Docker](https://www.docker.com);

## Como excutar

Inicie os contêineres usando o comando docker::

`docker-compose up -d`

Execute a migração do banco de dados:

`make migrate`

Para executar o projeto, execute o comando:

`cd cmd/ordersystem/ && go run main.go wire_gen.go`


## Servidor Web
Para testar o servidor web, execute os arquivos da pasta `/api` do projeto que são:

- `create_order.http`
- `list_orders.http`

## Servidor GraphQL

Para testar o GraphQL, acesse a url http://localhost:8080 e execute os comandos para criar e listar pedidos:
```graphql
mutation createOrder {
  createOrder(input:{id:"A", Price:30, Tax:2}) {
    id
    Price
    Tax
    FinalPrice
  }
}

query listOrders {
  orders {
    id
    Tax
    Price
    FinalPrice
  }
}
```

## Servidor gRPC

Para testar o gRPC, execute o comando com a ajuda de `evans` ([veja mais aqui](https://evans.syfm.me/)):

```shell
evans -r repl
package pb
service OrderService

```
Em seguida, use o comando `call CreateOrder` para criar um pedido e `call ListOrders` para listar os pedidos.