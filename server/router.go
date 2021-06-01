package server

import (
	"github.com/mrchar/seedpod/server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userHandler *handler.UserHandler
)

func newRouter() http.Handler {
	router := gin.Default()
	{
		userGroup := router.Group("/user")
		{
			userHandler = handler.DefaultUserHandler()
			userGroup.POST("/register", userHandler.Register)
			userGroup.POST("/login", userHandler.Login)
		}
	}
	return router
}
