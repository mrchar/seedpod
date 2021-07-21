package application

import (
	"github.com/mrchar/seedpod/account"
	"github.com/smartystreets/goconvey/convey"
	"testing"

	testUtils "github.com/mrchar/seedpod/utils/test"

	"github.com/sirupsen/logrus"
)

func init() {
	err := testUtils.DropTables()
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestMigrate(t *testing.T) {
	convey.Convey("表没有创建的情况下，创建表，不会返回错误", t, func() {
		err := testUtils.AutoMigrate(new(Application), new(Account), new(account.Account))
		convey.So(err, convey.ShouldBeNil)
	})

}

func TestManager_GetApplicationByAppId(t *testing.T) {
	convey.Convey("在数据库连接的情况下，创建Manager，不会返回错误", t, func() {
		err := testUtils.DropTables()
		convey.So(err, convey.ShouldBeNil)
		manager := DefaultManager()
		convey.Convey("在manager成功创建的情况下，执行AutoMigrate，不会返回错误", func() {
			err := testUtils.AutoMigrate(new(account.Account), new(Application), new(Account))
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("在Application不存在的情况下，创建Application，不会返回错误", func() {
				app, err := manager.AddApplication(Application{Name: "app"})
				convey.So(err, convey.ShouldBeNil)

				convey.Convey("在Application存在的情况下查询Application，不会返回错误", func() {
					app, err := manager.GetApplicationByAppId(app.AppId)
					convey.So(err, convey.ShouldBeNil)
					convey.Printf("Application: %+v", app)
				})
			})

		})
	})
}
