package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/account"
	"net/http"
)

type UserHandler struct {
	manager *account.Manager
}

func DefaultUserHandler() *UserHandler {
	manager := account.DefaultManager()
	return NewUserHandler(manager)
}

func NewUserHandler(manager *account.Manager) *UserHandler {
	return &UserHandler{
		manager: manager,
	}
}

type RegisterRequest struct {
	AccountName string `json:"accountName" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func (r RegisterRequest) ToJSON() []byte {
	bytes, _ := json.Marshal(r)
	return bytes
}

type RegisterResponse struct {
	Message string `json:"message,omitempty"`
}

func (u *UserHandler) Register(c *gin.Context) {
	var param RegisterRequest
	if err := c.BindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := u.manager.Register(param.AccountName, param.Password); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{"ok"})
}

type LoginRequest struct {
	AccountName string `json:"accountName" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message,omitempty"`
}

func (u *UserHandler) Login(c *gin.Context) {
	var param LoginRequest
	if err := c.BindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := u.manager.Login(param.AccountName, param.Password); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{"ok"})
}
