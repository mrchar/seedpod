package account

import (
	"encoding/base64"
	"github.com/mrchar/seedpod/common/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Account 描述账户
type Account struct {
	db.Model
	Name     string `gorm:"uniqueIndex"`
	Password string
	Roles    []*Role `gorm:"many2many:account_role"`
}

// BeforeSave 将Account保存到数据库之前加密Password
func (a *Account) BeforeSave(tx *gorm.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.Password = base64.RawStdEncoding.EncodeToString(hashed)
	return nil
}
