package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anucha-tk/go_bank/api"
	db "github.com/anucha-tk/go_bank/db/sqlc"
	"github.com/anucha-tk/go_bank/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	address  = "0.0.0.0:8080"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database_user := config.DBUserName
	database_password := config.DBPassword
	database_name := config.DBName
	database_host := config.DBHost

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
