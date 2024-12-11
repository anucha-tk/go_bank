package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database_user := os.Getenv("DATABASE_USERNAME")
	database_password := os.Getenv("DATABASE_PASSWORD")
	database_name := os.Getenv("DATABASE_NAME")
	database_host := os.Getenv("DATABASE_HOST")

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", database_user, database_password, database_host, database_name)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	// NOTE: best practice use conn.Ping() to check database
	if err := conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
