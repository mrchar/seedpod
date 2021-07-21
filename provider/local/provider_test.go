package local

import (
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/db"
	testUtils "github.com/mrchar/seedpod/utils/test"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRegisterAndLogin(t *testing.T) {
	convey.Convey("在数据库连接的情况下，创建Manager，不会返回错误", t, func() {
		err := testUtils.DropTables()
		convey.So(err, convey.ShouldBeNil)
		provider := Default()
		convey.Convey("在manager成功创建的情况下，执行AutoMigrate，不会返回错误", func() {
			err := db.Default().AutoMigrate(new(account.Account), new(account.Role), new(account.User))
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("在用户不存在的情况下，登录，会返回错误", func() {
				credential, err := provider.Login(NameAndPasswordCredential{Name: "account", Password: "password"})
				convey.Println(err)
				convey.So(err, convey.ShouldNotBeNil)
				convey.So(credential, convey.ShouldBeNil)
			})

			convey.Convey("在manager创建的情况下，注册账号，不会返回错误", func() {
				err := provider.Register(NameAndPasswordCredential{Name: "account", Password: "password"})
				convey.So(err, convey.ShouldBeNil)

				convey.Convey("在账户已经存在的情况下，再次注册账户，会返回错误", func() {
					err := provider.Register(NameAndPasswordCredential{Name: "account", Password: "password"})
					convey.Println(err)
					convey.So(err, convey.ShouldNotBeNil)
				})

				convey.Convey("在用户已经存在的情况下，登录，不会返回错误", func() {
					credential, err := provider.Login(NameAndPasswordCredential{Name: "account", Password: "password"})
					convey.So(err, convey.ShouldBeNil)
					convey.So(credential, convey.ShouldNotBeNil)
					convey.Printf("Token: %s", credential.Token)
				})
			})

		})
	})
}
