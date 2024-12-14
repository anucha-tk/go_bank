package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/anucha-tk/go_bank/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database_user := config.DBUserName
	database_password := config.DBPassword
	database_name := config.DBName
	database_host := config.DBHost

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", database_user, database_password, database_host, database_name)

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	// NOTE: best practice use conn.Ping() to check database
	if err := testDB.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
