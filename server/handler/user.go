package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/authenticator"
	"github.com/mrchar/seedpod/provider/local"
	"net/http"
)

var defaultUserHandler *UserHandler

type UserHandler struct {
	accountManager        *account.Manager
	authenticationManager *authenticator.Authenticator
	localProvider         *local.Provider
}

func DefaultUserHandler() *UserHandler {
	if defaultUserHandler == nil {
		manager := account.DefaultManager()
		authenticationManager := authenticator.DefaultAuthenticator()
		defaultUserHandler = NewUserHandler(manager, authenticationManager)
	}

	return defaultUserHandler
}

func NewUserHandler(accountManager *account.Manager, authenticationManager *authenticator.Authenticator) *UserHandler {
	return &UserHandler{
		accountManager:        accountManager,
		authenticationManager: authenticationManager,
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

// Register 注册账户
func (u *UserHandler) Register(c *gin.Context) {
	var param RegisterRequest
	if err := c.BindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := u.localProvider.Register(local.NameAndPasswordCredential{Name: param.AccountName, Password: param.Password}); err != nil {
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
	Token   string `json:"token"`
	Message string `json:"message,omitempty"`
}

func (u *UserHandler) Login(c *gin.Context) {
	var param LoginRequest
	if err := c.BindJSON(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	credential, err := u.localProvider.Login(local.NameAndPasswordCredential{Name: param.AccountName, Password: param.Password})
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: credential.Token})
}

type AuthRequest struct {
	AppId string `form:"appId" binding:"required"`
}

type AuthResponse struct {
	Token   string `json:"token"`
	Message string `json:"message,omitempty"`
}

func (u *UserHandler) Auth(c *gin.Context) {
	var param AuthRequest
	if err := c.BindQuery(&param); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	name := c.MustGet("token").(*local.LoginClaims).Name

	token, err := u.authenticationManager.Auth(name, param.AppId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, AuthResponse{Token: token})
}
