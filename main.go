package main

import (
	"database/sql"
	"github.com/anhuet/simplebank/api"
	db "github.com/anhuet/simplebank/db/sqlc"
	_ "github.com/lib/pq" // <------------ here
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://andy:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
