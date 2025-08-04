user-migrate-up:
	migrate -database "postgres://dkrnd:dkrnd@localhost:5433/sayur-user-service?sslmode=disable" -path ./user-service/database/migrations up

start:
	docker compose up -d

stop:
	docker compose down