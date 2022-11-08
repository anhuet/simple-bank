package api

import (
	db "github.com/anhuet/simplebank/db/sqlc"
	"github.com/anhuet/simplebank/token"
	"github.com/anhuet/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
	token  token.Maker
}

func NewServer(store *db.Store) *Server {
	config, err := util.LoadConfigFile(".")
	if err != nil {
		log.Fatal("Can not load a config", err)
	}
	tokenMaker, err := token.NewJwtMaker(config.TokenSecretKey)

	server := &Server{store: store, token: tokenMaker}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

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
