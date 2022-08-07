postgres:
	docker run --name policedb -p 5002:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it policedb createdb --username=root --owner=root police_db

dropdb:
	docker exec -it policedb dropdb --username=root --owner=root police_db

migrateup:
	migrate -path db/migration -database "postgresql://policeadmin:pcheck@123@policecheckserver.postgres.database.azure.com/police_db?sslmode=require" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://policeadmin:pcheck@123@policecheckserver.postgres.database.azure.com/police_db?sslmode=require" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc