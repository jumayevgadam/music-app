## run application
.PHONY: run
run:
	go run cmd/app/main.go

## migration_create is
.PHONY: migration_create
migration_create:
	migrate create -ext sql -dir ./internal/migrations/ -seq songs

## migration_up is
.PHONY: migration_up
migration_up:
	migrate -path ./internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/music_app?sslmode=disable" -verbose up

## migration_down is
.PHONY: migration_down
migration_down:
	migrate -path ./internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/music_app?sslmode=disable" -verbose down

## force_version is
.PHONY: force_version
force_version:
	migrate -path ./internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/music_app?sslmode=disable" force 1
