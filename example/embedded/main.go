package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/server/handler"
	"github.com/mrchar/seedpod/server/middleware"
	"net/http"
)

func main() {
	server := gin.Default()
	userGroup := server.Group("/api/user")
	{
		userHandler := handler.DefaultUserHandler()
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}

	authMidWare := middleware.DefaultAuthMiddleware()
	authorizedGroup := server.Group("/api/authorized", authMidWare.VerifyLogin)
	{
		authorizedGroup.GET("/greeting",
			func(context *gin.Context) {
				name := context.Param("name")
				if name == "" {
					name = "you"
				}
				context.JSON(http.StatusOK,
					gin.H{"message": fmt.Sprintf("hello %s", name)},
				)
			})
	}
}
