postgres:
	docker run -d -p 5436:5432 -e POSTGRES_USER=letmego -e POSTGRES_PASSWORD=letmego -e POSTGRES_DB=letmego -e PGDATA=/var/lib/postgresql/data/pgdata postgres:13

compile:
	CGO_ENABLED=0 go build -o build/let-me-go -ldflags="-s -w"

dev:
	go build && ENVIRONMENT=dev ./let-me-go

new_migration:
	migrate create -ext sql -dir ./migrations -seq $(MIGRATION_NAME)

migrate:
	#POSTGRES_URL=postgres://letmego:letmego@localhost:5436/letmego?sslmode=disable make migrate
	migrate -database $(POSTGRES_URL) -path migrations up


.PHONY: postgres compile dev new_migration migrate
