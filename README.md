# Go Transactions
This project is an API using the [Go programming language](https://golang.org/).

## Requires
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed.

## Nice to know
- [GoLang](https://golang.org/) version 1.14.10
- Web router [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [Clean Architecture Concept](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Project Layout](https://github.com/golang-standards/project-layout)

## Get started
The essential commands are provided from the [Makefile]().

To run the project tests you can use the command `test`. For example:
```
make test
```

So to run the project in localhost as a development mode you can use the command `start`
```
make start
```

The `start` will up the http server in your machine in [localhost:3000](http://localhost:3000)

### Other make commands
- `make build` - build the image docker
- `make lint` - runs Go linters
- `make bash` - open the container bash
- `make restart` - restart the docker compose and containers
- `make stop` - stop the containers
- `make clean` - stop containers and clean the environment
- `make logs [tail=50]` - show the logs (tail is optional, the default value is 100)

## How to use
This chapter show about how to use the API interface.
We'll use the [curl](https://curl.haxx.se/) tool to show the examples.

## Create a new account
Send a request to `/accounts` with the `POST` verb.

For example:
```bash
curl -X POST \
  localhost:3000/accounts \
  -d '{
    "account_id": 1,
    "document_number": "12345678900"
  }'
```

Response:
```json
{
  "account_id": "1",
  "document_number": "12345678900"
}
```

## Get account informations
Send a request to `/accounts/:accountId` with the `GET` verb, passing the `accountId` as account identifier.

For example:
```bash
curl localhost:3000/accounts/1
```

Response:
```json
{
  "account_id": 1,
  "document_number": "12345678900"
}
```

## Register a new transaction
Send a request to `/transactions` with the `POST` verb.

For example:
```bash
curl -X POST \
  localhost:3000/transactions \
  -d '{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": 123.45
  }'
```

## Get all transactions
Send a request to `/transactions` with the `GET` verb.

For example:
```bash
curl localhost:3000/transactions
```
