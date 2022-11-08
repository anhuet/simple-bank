package main

import (
	"database/sql"
	"github.com/anhuet/simplebank/api"
	db "github.com/anhuet/simplebank/db/sqlc"
	"github.com/anhuet/simplebank/util"
	_ "github.com/lib/pq" // <------------ here
	"log"
)

func main() {
	config, err := util.LoadConfigFile(".")
	if err != nil {
		log.Fatal("Can not load a config", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}
	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
