postgres:
	sudo docker run --name postgres -p 5433:5432 -e POSTGRES_USER=andy -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	sudo docker exec -it postgres createdb --username=andy --owner=andy simple_bank

dropdb:
	sudo docker exec -it postgres dropdb simle_bank

migrateup:
	migrate -path db/migration -database "postgresql://andy:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://andy:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://andy:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb sql migrateup migratedown sqlc migrateup1