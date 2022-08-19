package api

import (
	db "github.com/Nuwan-Walisundara/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store db.Store
	route *gin.Engine
}

func NewServer(store db.Store) *Server {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server := &Server{store: store}
	router := gin.Default()
	server.route = router
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/", server.listAccounts)
	router.POST("/transfer", server.createTransfer)
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error :": err.Error()}
}

func (server *Server) Start(address string) error {

	return server.route.Run(address)
}
