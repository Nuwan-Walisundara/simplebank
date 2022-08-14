package main

import (
	"database/sql"
	"log"

	"github.com/Nuwan-Walisundara/simplebank/api"
	db "github.com/Nuwan-Walisundara/simplebank/db/sqlc"
	"github.com/Nuwan-Walisundara/simplebank/util"
	_ "github.com/lib/pq"
)

 
func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Unable read the configurations ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.Data_Source)

	if err != nil {
		log.Fatal("Unable start the server ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.Server_Address)

	if err != nil {
		log.Fatal("Unable start the server ", err)
	}

}
