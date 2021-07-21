package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/application"
	"net/http"
)

var defaultApplicationHandler *ApplicationHandler

type ApplicationHandler struct {
	applicationManager *application.Manager
}

func DefaultApplicationHandler() *ApplicationHandler {
	if defaultUserHandler == nil {
		manager := application.DefaultManager()
		defaultApplicationHandler = NewApplicationHandler(manager)
	}
	return defaultApplicationHandler
}

func NewApplicationHandler(applicationManager *application.Manager) *ApplicationHandler {
	return &ApplicationHandler{
		applicationManager: applicationManager,
	}
}

type RegisterApplicationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RegisterApplicationResponse struct {
	AppId   string `json:"appId" binding:"required"`
	Message string `json:"message,omitempty"`
}

// RegisterApplication 注册应用程序
func (a *ApplicationHandler) RegisterApplication(c *gin.Context) {
	var param RegisterApplicationRequest
	if err := c.BindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	appId, err := a.applicationManager.RegisterApplication(param.Name, param.Description)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, RegisterApplicationResponse{AppId: appId, Message: "ok"})
}
