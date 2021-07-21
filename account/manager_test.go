package account

import (
	"testing"

	testUtils "github.com/mrchar/seedpod/utils/test"

	"github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
)

func init() {
	err := testUtils.DropTables()
	if err != nil {
		logrus.Fatal(err)
	}
}

// 测试创建的表关系是否正确
func TestMigrate(t *testing.T) {
	convey.Convey("在表不存在的情况下，执行AutoMigrate，不会返回错误", t, func() {
		err := testUtils.AutoMigrate(new(Account), new(Profile), new(MobilePhone), new(Email), new(Role), new(User))
		convey.So(err, convey.ShouldBeNil)
	})
}
