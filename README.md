# seedpod

seedpod是一个简易用户管理系统。

你有两种方式使用seedpod，并且都非常简单易用。

1. 将seedpod作为独立的服务部署，使用合适的方式调用seedpod。当您同时使用多个服务为同样的人提供服务时，您不必在各个服务中重复添加用户管理系统。
2. 直接引入seedpod核心库，将它作为你的程序的一部分。配置好合适的数据库，seedpod会自动创建用户管理需要的用户表和角色表等。您不需要再为您的程序专门开发用户管理。

## 快速开始

### 作为服务独立部署

#### Linux

```shell
docker-compose up
```

#### Windows

```shell
docker compose up
```

### 引入seedpod核心库

```go
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
```


