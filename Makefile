up:
	docker-compose up -d

down:
	docker-compose down

run/api:
	go run ./cmd/api

migration/create:
	migrate create -ext=.sql -dir=./data/migrations ${migration_name}

migration/up:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev up

migration/down-one:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev down 1

migration/down:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev down

migration/force:
	migrate -path=./data/migrations -database=mysql://realworld:realworld@/realworld_dev force ${version}
