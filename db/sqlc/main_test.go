package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:postgresroot@localhost:5432/simple_bank?sslmode=disable"
)

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Unable to connect to the db :", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
