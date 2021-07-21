package account

import "github.com/mrchar/seedpod/common/db"

// Profile 用于对账户进行描述
type Profile struct {
	db.Model
	AccountUUID            string
	Account                Account
	Name                   string
	PrimaryMobilePhoneUUID string
	PrimaryMobilePhone     MobilePhone
	MobilePhone            []MobilePhone
	PrimaryEmailUUID       string
	PrimaryEmail           Email
	Emails                 []Email
	PrimaryAddressUUID     string
	PrimaryAddress         Address
	Addresses              []Address
}

// MobilePhone 记录手机号码
type MobilePhone struct {
	db.Model
	ProfileUUID string
	Prefix      string // 国家地区码
	Number      string // 手机号码
}

type Email struct {
	db.Model
	ProfileUUID string
	Address     string // 电子邮件地址
}

type Address struct {
	db.Model
	ProfileUUID string
	Country     string
	Province    string
	City        string
	Detail      string
}
