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

## db/migration/create migration_name=$1: create a new database migration 
db/migration/create:
	migrate create -ext=.sql -dir=./data/migrations ${migration_name}

## db/migration/up: apply all up database migrations
db/migration/up:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev up

## db/migration/down-one: apply one down database migration
db/migration/down-one:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev down 1

## db/migration/down: apply all down database migrations
db/migration/down:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev down

## db/migration/force version=$1: force specific migration version
db/migration/force:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev force ${version}
