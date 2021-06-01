package account

import (
	testUtils "github.com/mrchar/seedpod/utils/test"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestManager(t *testing.T) {
	convey.Convey("在数据库连接的情况下，创建Manager，不会返回错误", t, func() {
		err := testUtils.DropTables()
		convey.So(err, convey.ShouldBeNil)
		manager := DefaultManager()
		convey.Convey("在manager成功创建的情况下，执行AutoMigrate，不会返回错误", func() {
			err := manager.AutoMigrate()
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("在用户不存在的情况下，登录，会返回错误", func() {
				err := manager.Login("account", "password")
				convey.Println(err)
				convey.So(err, convey.ShouldNotBeNil)
			})

			convey.Convey("在manager创建的情况下，注册账号，不会返回错误", func() {
				err := manager.Register("account", "password")
				convey.So(err, convey.ShouldBeNil)

				convey.Convey("在账户已经存在的情况下，再次注册账户，会返回错误", func() {
					err := manager.Register("account", "password")
					convey.Println(err)
					convey.So(err, convey.ShouldNotBeNil)
				})

				convey.Convey("在用户已经存在的情况下，登录，不会返回错误", func() {
					err := manager.Login("account", "password")
					convey.So(err, convey.ShouldBeNil)
				})
			})

		})
	})
}
