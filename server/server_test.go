package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/jwt"
	"github.com/mrchar/seedpod/server/middleware"
	testUtils "github.com/mrchar/seedpod/utils/test"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserHandler(t *testing.T) {
	convey.Convey("数据库启动的情况下，启动Web服务器，不会返回错误", t, func() {
		err := testUtils.DropTables()
		convey.So(err, convey.ShouldBeNil)
		err = account.DefaultManager().AutoMigrate()
		convey.So(err, convey.ShouldBeNil)

		mockServer := newRouter()
		convey.Convey("用户不存在的情况下，登录，会返回错误", func() {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(
				http.MethodPost,
				"/user/login",
				bytes.NewBufferString(`{"accountName":"account", "password":"password"}`),
			)
			mockServer.ServeHTTP(w, r)
			convey.So(w.Code, convey.ShouldNotEqual, http.StatusOK)
		})

		convey.Convey("用户不存在的情况下，注册用户，不会返回错误", func() {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(
				http.MethodPost,
				"/user/register",
				bytes.NewBufferString(`{"accountName":"account", "password":"password"}`),
			)
			mockServer.ServeHTTP(w, r)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)

			convey.Convey("用户已经存在的情况下，注册用户，会返回错误", func() {
				w := httptest.NewRecorder()
				r, _ := http.NewRequest(
					http.MethodPost,
					"/user/register",
					bytes.NewBufferString(`{"accountName":"account", "password":"password"}`),
				)
				mockServer.ServeHTTP(w, r)
				convey.So(w.Code, convey.ShouldNotEqual, http.StatusOK)
			})

			convey.Convey("用户存在的情况下，登录，不会返回错误", func() {
				w := httptest.NewRecorder()
				r, _ := http.NewRequest(
					http.MethodPost,
					"/user/login",
					bytes.NewBufferString(`{"accountName":"account", "password":"password"}`),
				)
				mockServer.ServeHTTP(w, r)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.Printf("Resp: %s", w.Body.String())
			})
		})
	})
}

func TestAuthMiddleware(t *testing.T) {
	convey.Convey("数据库启动的情况下，启动Web服务器，不会返回错误", t, func() {
		router := gin.Default()
		router.GET("/hello/world", middleware.DefaultAuthMiddleware().VerifyLogin, func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World")
		})

		convey.Convey("没有登录的情况下，访问/hello/world， 会返回401", func() {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(
				http.MethodGet,
				"/hello/world",
				nil,
			)
			router.ServeHTTP(w, r)
			convey.So(w.Code, convey.ShouldEqual, http.StatusUnauthorized)
		})
		convey.Convey("登录的情况下，访问/hello/world，会返回200", func() {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(
				http.MethodGet,
				"/hello/world",
				nil,
			)
			token, err := jwt.Default().Issue(&account.LoginClaims{
				Issuer:    "seedpod",
				Name:      "hello",
				Roles:     nil,
				ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			})
			convey.So(err, convey.ShouldBeNil)

			r.Header.Add("Authorization", "Bearer "+token)

			router.ServeHTTP(w, r)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		})
	})
}
