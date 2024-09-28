# Serviço de Pedidos

Este projeto implementa um serviço de pedidos com endpoints REST, gRPC e GraphQL para listar pedidos.

## Pré-requisitos

- Docker e Docker Compose
- Go 1.20 ou posterior (para desenvolvimento local)

## Como começar

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/projeto-pedidos.git
   cd projeto-pedidos
   ```

2. Inicie os serviços usando Docker Compose:
   ```
   docker-compose up --build
   ```

3. Os serviços estarão disponíveis em:
   - API REST: http://localhost:8080/order
   - gRPC: localhost:50051
   - GraphQL: http://localhost:8080/graphql

## Endpoints da API

### REST

- Listar Pedidos: GET http://localhost:8080/order

### gRPC

O serviço gRPC está disponível na porta 50051. Você pode usar um cliente gRPC para interagir com ele.

### GraphQL

- GraphQL Playground: http://localhost:8080/graphql
- Consulta para listar pedidos:
  ```graphql
  query {
    listarPedidos {
      id
      nomeCliente
      valorTotal
      status
      criadoEm
    }
  }
  ```

## Desenvolvimento

Para desenvolvimento local sem Docker:

1. Instale as dependências:
   ```
   go mod download
   ```

2. Execute as migrações:
   ```
   migrate -database "postgresql://postgres:postgres@localhost:5432/pedidos?sslmode=disable" -path internal/database/migrations up
   ```

3. Execute a aplicação:
   ```
   go run cmd/server/main.go
   ```

## Testando

Você pode usar o arquivo `api.http` para testar os endpoints REST e GraphQL usando um cliente HTTP que suporte arquivos `.http` (como a extensão REST Client do VS Code).

## Portas

- Servidor HTTP (REST e GraphQL): 8080
- Servidor gRPC: 50051