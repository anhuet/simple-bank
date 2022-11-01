package api

import (
	db "github.com/anhuet/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account", server.ListAccounts)
	router.GET("/account/:id", server.GetAccount)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}
