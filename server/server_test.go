package server

import (
	"bytes"
	"github.com/mrchar/seedpod/account"
	testUtils "github.com/mrchar/seedpod/utils/test"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
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
