package account

import "github.com/mrchar/seedpod/common/db"

// User 对帐户的拥有者的描述
type User struct {
	db.Model
	AccountUUID string
	Account     Account
	Name        string
	Gender      string
	Country     string
	Province    string
	City        string
}
