package authenticator

import (
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/application"
	testUtils "github.com/mrchar/seedpod/utils/test"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAuth(t *testing.T) {
	convey.Convey("在数据库连接的情况下，创建Manager，不会返回错误", t, func() {
		err := testUtils.DropTables()
		convey.So(err, convey.ShouldBeNil)
		manager := DefaultAuthenticator()
		convey.Convey("在manager成功创建的情况下，执行AutoMigrate，不会返回错误", func() {
			err := testUtils.AutoMigrate(new(account.Account), new(application.Application), new(application.Account))
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("在账户不存在情况下，注册账号，不会返回错误", func() {
				err := manager.accountManager.RegisterAccount("account", "password")
				convey.So(err, convey.ShouldBeNil)

				convey.Convey("在应用不存在的情况下，注册应用，不会返回错误", func() {
					appId, err := manager.applicationManager.RegisterApplication("app", "description")
					convey.So(err, convey.ShouldBeNil)

					convey.Convey("在用户和应用存在的情况下，验证用户，不会返回错误", func() {
						token, err := manager.Auth("account", appId)
						convey.So(err, convey.ShouldBeNil)
						convey.So(token, convey.ShouldNotBeEmpty)
						convey.Printf("Token: %s", token)
					})
				})

			})

		})
	})
}
