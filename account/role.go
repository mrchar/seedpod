package account

import "github.com/mrchar/seedpod/common/db"

type Role struct {
	db.Model
	Name     string     `gorm:"column:name;uniqueIndex"`
	Accounts []*Account `gorm:"many2many:account_role"`
}
