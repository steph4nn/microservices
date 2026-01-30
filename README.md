# Como rodar (somente docker compose)

Este projeto sobe 4 containers:

- `mysql` (MySQL 8.0)
- `payment` (gRPC na porta 3001)
- `shipping` (gRPC na porta 3002)
- `order` (gRPC na porta 3000)

O arquivo [microservices/init.sql](init.sql) cria os bancos `order`, `payment` e `shipping` automaticamente no primeiro start do MySQL.

## Subir tudo

Execute a partir da **raiz do repo**:

- `docker compose -f microservices/docker-compose.yml up --build`

## Parar

- `docker compose -f microservices/docker-compose.yml down`

## Reset (apagar dados do MySQL)

Isso remove o volume persistente do banco.

- `docker compose -f microservices/docker-compose.yml down -v`

## Ver logs

- logs de tudo:
  - `docker compose -f microservices/docker-compose.yml logs -f`

- logs de um servico:
  - `docker compose -f microservices/docker-compose.yml logs -f payment`
  - `docker compose -f microservices/docker-compose.yml logs -f shipping`
  - `docker compose -f microservices/docker-compose.yml logs -f order`

## Banco (MySQL)

Credenciais (definidas no compose):

- usuario: `root`
- senha: `minhasenha`
