package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //blank identifier is as this driver is being used without calling any function
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource) //to create a db connection
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB) //Initialize a Queries struct with "conn" where "conn" implements DBTX interface in db.go
	os.Exit(m.Run())          //to start running unit tests
}
