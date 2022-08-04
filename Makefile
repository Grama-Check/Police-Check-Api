postgres:
	docker run --name policedb -p 5002:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it policedb createdb --username=root --owner=root police_db

dropdb:
	docker exec -it policedb dropdb --username=root --owner=root police_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5002/police_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5002/police_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc