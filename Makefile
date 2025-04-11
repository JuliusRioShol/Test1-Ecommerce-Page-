## Filename Makefile
include .envrc

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet 
vet: fmt 
	go vet ./...

.PHONY: run 
run: vet
	go run ./cmd/ -addr=":4000" -dsn=${ECOMERCE_DB_DSN}

.PHONY: db/psql
db/psql:
	psql ${ECOMERCE_DB_DSN}

.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up 
db/migrations/up :
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${ECOMERCE_DB_DSN} up 

.PHONY:db/migrations/force 
db/migrations/force:
	@echo 'Forcing migrations because dirty'
	migrate  -path=./migrations -database=${ECOMERCE_DB_DSN} force ${number}

.PHONY: db/migrations/rollback
db/migrations/rollback:
	@echo 'Rolling back to succesfull migrations: ${version}'
	migrate  -path=./migrations -database=${ECOMERCE_DB_DSN} down ${version}