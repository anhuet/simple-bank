package api

import (
	"fmt"
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
	config util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	config, err := util.LoadConfigFile(".")
	if err != nil {
		log.Fatal("Can not load a config", err)
	}
	tokenSecret, err := token.NewJWTMaker(config.TokenSecretKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server := &Server{store: store, token: tokenSecret}
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.token))

	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/account", server.ListAccounts)
	authRoutes.GET("/account/:id", server.GetAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

}
