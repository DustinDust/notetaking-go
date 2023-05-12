postgres:
	docker run --name postgresql_db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123123 -d postgres

createdb:
	docker exec -it postgresql_db createdb --username=postgres --owner=postgres notetaking-go

dropdb:
	docker exec -it postgresql_db dropdb --username=postgres notetaking-go

migrateup:
	migrate -path migration -database "postgres://postgres:123123@localhost:5432/notetaking-go?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgres://postgres:123123@localhost:5432/notetaking-go?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown