# Objetivo do Projeto
Este projeto tem como objetivo implementar a funcionalidade de criação e listagem de orders utilizando três diferentes abordagens de comunicação:

- REST HTTP: Expor endpoints REST para criação (POST /orders) e listagem (GET /list-orders) de orders.
- gRPC: Implementar um serviço gRPC com métodos para criar e listar orders.
- GraphQL: Disponibilizar queries e mutations para criação e listagem de orders através de uma API GraphQL.

# Como rodar o projeto
Para executar o projeto localmente com todos os serviços necessários (banco de dados, RabbitMQ e sua aplicação), basta rodar o comando:

```
docker-compose up -d --build
```

Isso irá:

 1. Buildar as imagens Docker (caso necessário)
 2. Subir os containers com MySQL, RabbitMQ e sua aplicação
 3. A aplicação estará disponível nas seguintes portas:
    - REST API: Porta 8000
    - gRPC Service: Porta 50051
    - GraphQL Server: Porta 8080

# Testando
REST HTTP:
Utilize o arquivo orders.http para testar as requisições de criação e listagem de orders.

gRPC:
Use o Evans para testar o serviço gRPC. Lembre-se de selecionar o pacote pb e o serviço OrderService.

GraphQL:
Acesse a interface gráfica disponível na porta 8080 para realizar queries e mutations.