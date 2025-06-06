## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## up: start development environment using docker
up:
	docker-compose up -d

## down: stop development environment using docker
down:
	docker-compose down

## run/api: run the cmd/api application
run/api:
	go run ./cmd/api

## test: run app unit tests
test:
	go test ./...

## db/migration/create migration_name=$1: create a new database migration 
db/migration/create:
	migrate create -ext=.sql -dir=./migrations ${migration_name}

## db/migration/up: apply all up database migrations
db/migration/up:
	migrate -path=./migrations -database=mysql://realworld:realworld@/realworld_dev up

## db/migration/down-one: apply one down database migration
db/migration/down-one:
	migrate -path=./migrations -database=mysql://realworld:realworld@/realworld_dev down 1

## db/migration/down: apply all down database migrations
db/migration/down:
	migrate -path=./migrations -database=mysql://realworld:realworld@/realworld_dev down

## db/migration/force version=$1: force specific migration version
db/migration/force:
	migrate -path=./migrations -database=mysql://realworld:realworld@/realworld_dev force ${version}

## mock/generate: generate mocks
mock/generate:
	docker run -v "$PWD":/src -w /src vektra/mockery --all
