package account

import "github.com/mrchar/seedpod/common/db"

// Role 分配给账户的角色
type Role struct {
	db.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Accounts    []*Account `gorm:"many2many:account_role"`
}
