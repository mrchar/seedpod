package server

import (
	"github.com/mrchar/seedpod/server/handler"
	"github.com/mrchar/seedpod/server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userHandler        = handler.DefaultUserHandler()
	applicationHandler = handler.DefaultApplicationHandler()
)

func newRouter() http.Handler {
	router := gin.Default()
	{
		userGroup := router.Group("/user")
		{
			userGroup.POST("/register", userHandler.Register)
			userGroup.POST("/login", userHandler.Login)
		}
		// 需要登录才能调用的接口
		authorizedGroup := router.Group("", middleware.DefaultAuthMiddleware().VerifyLogin)
		{
			userGroup := authorizedGroup.Group("/user")
			{
				userGroup.POST("auth", userHandler.Auth)
			}
			appGroup := authorizedGroup.Group("/app")
			{
				appGroup.POST("/add", applicationHandler.RegisterApplication)
			}
		}
	}
	return router
}
