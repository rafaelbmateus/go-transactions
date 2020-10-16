include scripts/help.mk

version ?= latest
image = rafaelbmateus/go-transactions:$(version)
docker = docker run --rm $(image)
compose = docker-compose

.PHONY: build
build: ##@development build image docker base
	docker build -f build/Dockerfile . -t $(image)

.PHONY: lint
lint: build ##@lint static analysis code base
	$(docker) golangci-lint run

.PHONY: test
test: build ##@test run all tests
	$(docker) go test ./...

.PHONY: bash
bash: ##@development Start a bash session within the container
	$(compose) run --rm app /bin/bash

.PHONY: start
start: ##@development start development environment in background
	$(compose) up --build -d

.PHONY: restart
restart: ##@development Restart development environment [service="svc1 svc2..."]
	$(compose) restart

.PHONY: stop
stop: ##@test Stop development environment
	$(compose) down -v --remove-orphans

.PHONY: clean
clean: ##@development Clean the development environment
	rm -rf public;
	$(compose) down \
		--remove-orphans \
		--volumes

tail ?= 100
.PHONY: logs
logs: ##@development Follow development logs
	$(compose) logs -f --tail=$(tail) $(service)
