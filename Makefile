postgres:
	docker run -d --name postgres13 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres

createdb:
	docker exec -it postgres13 createdb --username=postgres --owner=postgres vedo-testovi

dropdb:
	docker exec -it postgres13 dropdb vedo-testovi -U postgres

migrateup:
	migrate -path db -database "postgresql://postgres:postgres@localhost:5432/vedo-testovi?sslmode=disable" -verbose up

migratedown:
	migrate -path db -database "postgresql://postgres:postgres@localhost:5432/vedo-testovi?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
