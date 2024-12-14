package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/anucha-tk/go_bank/api"
	db "github.com/anucha-tk/go_bank/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	address  = "0.0.0.0:8080"
)

func main() {
	err := godotenv.Load()
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

	if err := conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
