package account

import (
	"encoding/base64"
	"github.com/mrchar/seedpod/db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var defaultManager *Manager

// Manager 用户管理Account
type Manager struct {
	db *gorm.DB
}

// DefaultManager 使用默认的数据库创建Manager
func DefaultManager() *Manager {
	if defaultManager == nil {
		defaultManager = NewManager(db.Default())
	}
	return defaultManager
}

// NewManager 新建Manager
func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) AutoMigrate() error {
	return m.db.AutoMigrate(new(Account))
}

// Register 注册账户
func (m *Manager) Register(name, password string) error {
	if m.accountExists(name) {
		return errors.New("同名账户已经存在")
	}

	_, err := m.addAccount(name, password)
	if err != nil {
		return errors.Wrap(err, "持久化账户时发生错误")
	}
	return nil
}

// Login 登录账户
func (m *Manager) Login(name, password string) error {
	account, err := m.getAccount(name)
	if err != nil {
		return errors.Wrap(err, "用户名或密码错误")
	}

	hashedPassword, err := base64.RawStdEncoding.DecodeString(account.Password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return errors.Wrap(err, "用户名或密码错误")
	}
	return nil
}

func (m *Manager) addAccount(name, password string) (*Account, error) {
	account := Account{
		Name:     name,
		Password: password,
	}

	result := m.db.Create(&account)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (m *Manager) getAccount(name string) (*Account, error) {
	var account Account
	result := m.db.First(&account, "name=?", name)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

func (m *Manager) accountExists(name string) bool {
	var count int64
	if m.db.Where("name = ?", name).Count(&count); count > 0 {
		return true
	}
	return false
}

func (m *Manager) addRole(name string) (*Role, error) {
	role := Role{
		Name: name,
	}

	result := m.db.Create(&role)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &role, nil
}
