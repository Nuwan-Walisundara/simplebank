package api

import (
	db "github.com/Nuwan-Walisundara/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.route = router
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/", server.listAccounts)
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error :": err.Error()}
}

func (server *Server) Start(address string) error {

	return server.route.Run(address)
}
