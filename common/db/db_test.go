package db

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefault(t *testing.T) {
	convey.Convey("数据库运行的情况下，使用配置文件初始化数据库连接，不会返回错误", t, func() {
		Default()
	})
}
