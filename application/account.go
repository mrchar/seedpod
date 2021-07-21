package application

import (
	"crypto/rand"
	"fmt"
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Account 应用程序使用的账户
type Account struct {
	db.Model
	OpenId          string `gorm:"column:open_id"` // seedpod 作为provider提供账户信息时使用的openId，用于代替account.Account.UUID
	ApplicationUUID string `gorm:"column:application_uuid"`
	Application     *Application
	AccountUUID     string `gorm:"column:account_uuid"`
	Account         *account.Account
}

func (a *Account) TableName() string {
	return "application_account"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if err := a.Model.BeforeCreate(tx); err != nil {
		return err
	}
	// 创建openId
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		err = errors.Wrap(err, "创建AppId时发生错误")
		return err
	}

	a.OpenId = fmt.Sprintf("op%x", buffer)
	return nil
}
