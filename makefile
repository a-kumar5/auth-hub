postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=admin -p 5432:5432 -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres authdb

dropdb:
	docker exec -it postgres dropdb --username=postgres authdb

migrateup:
	migrate -path db/migration -database "postgres://postgres:admin@localhost:5433/authdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:admin@localhost:5433/authdb?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown