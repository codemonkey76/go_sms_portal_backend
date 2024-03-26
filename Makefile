include .env
export

DATABASE_URL="postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate-up:
	migrate -path db/migration -database ${DATABASE_URL} -verbose up

migrate-down:
	migrate -path db/migration -database ${DATABASE_URL} -verbose down

migrate-reset:
	docker-compose down
	docker-compose up -d
	sleep 2
	migrate -path db/migration -database ${DATABASE_URL} -verbose up

