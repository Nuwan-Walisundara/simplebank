DB_URL=postgresql://postgres:postgresroot@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
		
test:
	go test -v -cover ./...
sqlc:
	sqlc generate	
server:
	go run main.go

.PHONY:	migrateup migratedown test  sqlc server