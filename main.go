package main

import (
	"database/sql"
	"log"

	"github.com/Nuwan-Walisundara/simplebank/api"
	db "github.com/Nuwan-Walisundara/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:postgresroot@localhost:5432/simple_bank?sslmode=disable"
	address        = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal("Unable start the server ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(address)

	if err != nil {
		log.Fatal("Unable start the server ", err)
	}

}
