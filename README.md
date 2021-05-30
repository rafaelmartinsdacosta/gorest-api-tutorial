# GO REST tutorial
Tutorial golang para criação de APIs REST usando Mux.

### Docker-compose
Iniciando banco de dados + aplicação REST API

```
docker-compose up --build
```

### :elephant: Iniciando container PostgreSQL
Criar banco de dados

```
docker run --name comments-api-db -e POSTGRES_PASSWORD=postgres -p 54321:5432 -d postgres
```

### :key: Variáveis de ambiente
Revisar o arquivo env.sh conforme a necessidade

```
. env.sh
```

### :fire: Teste unitários 
Executando os testes de unidade
```
go test ./... -tags=e2e -v
```