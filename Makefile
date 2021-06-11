ifneq (,$(wildcard ./.env))
	include .env
	export
endif

postgres:
	docker container run -dt --name postgres -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} postgres:12-alpine

createdb:
	docker container exec -it postgres createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker container exec -it postgres dropdb --username=${DB_USER} ${DB_NAME}

startdb:
	docker container start postgres

stopdb:
	docker container stop postgres

migrateup:
	migrate -path db/migrations -database "${DB_DIALECT}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose up

migratedown:
	migrate -path db/migrations -database "${DB_DIALECT}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose down

sqlc:
	sqlc generate

run:
	nodemon --exec go run main.go --signal SIGTERM

.PHONY: postgres createdb dropdb startdb stopdb migrateup migratedown sqlc run
