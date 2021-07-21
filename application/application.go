package application

import (
	"crypto/rand"
	"fmt"
	"github.com/mrchar/seedpod/common/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Application 注册的seedpod中的应用程序
type Application struct {
	db.Model
	AppId       string `gorm:"column:app_id;uniqueIndex"`
	AppSecret   string `gorm:"column:app_secret"`
	Name        string `gorm:"column:name;uniqueIndex"`
	Description string `gorm:"column:description"`
	Accounts    []Account
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	if err := a.Model.BeforeCreate(tx); err != nil {
		return err
	}

	// 创建appId
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		err = errors.Wrap(err, "创建AppId时发生错误")
		return err
	}

	a.AppId = fmt.Sprintf("ap%x", buffer)
	return nil
}
