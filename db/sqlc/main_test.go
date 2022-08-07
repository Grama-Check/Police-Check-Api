package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://policeadmin:pcheck@123@policecheckserver.postgres.database.azure.com/police_db?sslmode=require"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db : ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
