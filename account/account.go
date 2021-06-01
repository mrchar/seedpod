package account

import (
	"encoding/base64"
	"github.com/mrchar/seedpod/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Account 描述账户
type Account struct {
	db.Model
	Name     string  `gorm:"column:name;uniqueIndex"`
	Password string  `gorm:"column:password"`
	Roles    []*Role `gorm:"many2many:account_role"`
}

func (a *Account) BeforeSave(tx *gorm.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.Password = base64.RawStdEncoding.EncodeToString(hashed)
	return nil
}
