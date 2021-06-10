package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/jwt"
	"net/http"
	"strings"
)

var defaultAuthMiddleware *AuthMiddleware

type AuthenticationResponse struct {
	Message string `json:"message"`
}

type AuthMiddleware struct {
	iver *jwt.IssueVerifier
}

func DefaultAuthMiddleware() *AuthMiddleware {
	if defaultAuthMiddleware == nil {
		defaultAuthMiddleware = NewAuthMiddleware(jwt.Default())
	}

	return defaultAuthMiddleware
}

func NewAuthMiddleware(iver *jwt.IssueVerifier) *AuthMiddleware {
	return &AuthMiddleware{
		iver: iver,
	}
}

func (a *AuthMiddleware) VerifyLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	split := strings.Split(token, " ")
	if token == "" || !strings.HasPrefix(token, "Bearer") || len(split) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, AuthenticationResponse{"没有找到token"})
		return
	}

	token = strings.TrimSpace(split[1])
	claims := &account.LoginClaims{}
	t, err := a.iver.Verify(token, claims)
	if err != nil {
		// TODO: 处理具体错误
		c.AbortWithStatusJSON(http.StatusUnauthorized, AuthenticationResponse{err.Error()})
		return
	}

	claims, ok := t.Claims.(*account.LoginClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, AuthenticationResponse{"使用了错误的令牌"})
		return
	}

	c.Set("token", claims)
	c.Next()
}
