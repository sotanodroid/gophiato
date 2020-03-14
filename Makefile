.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: run
run:
	./apiserver -config-path configs/apiserver.toml

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate_up
migrate_up:
	 ./migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable" up

.PHONY: migrate_down
migrate_down:
	 ./migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable" down

.DEFAULT_GOAL := build
