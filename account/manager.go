package account

import (
	"github.com/mrchar/seedpod/common/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var defaultManager *Manager

// Manager 管理Account的Manager
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

// RegisterAccount 注册账户
func (m *Manager) RegisterAccount(name, password string) error {
	// 检查名称是否可用
	if m.AccountExists(name) {
		return errors.New("同名账户已经存在")
	}

	account := Account{
		Name:     name,
		Password: password,
	}

	// 添加账户
	_, err := m.AddAccount(account)
	if err != nil {
		return errors.Wrap(err, "持久化账户时发生错误")
	}
	return nil
}

// AddAccount 添加账户
func (m *Manager) AddAccount(account Account) (*Account, error) {
	result := m.db.Create(&account)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &account, nil
}

// ListAllAccounts 获取所有的用户
// TODO: 分页
func (m *Manager) ListAllAccounts() ([]Account, error) {
	var accounts []Account
	result := m.db.Find(&accounts)
	if err := result.Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccount 根据账户名获取账户
func (m *Manager) GetAccount(name string) (*Account, error) {
	var account Account
	result := m.db.First(&account, "name=?", name)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

// AccountExists 检查账户名对应的账户是否存在
func (m *Manager) AccountExists(name string) bool {
	var count int64
	if m.db.Model(new(Account)).Where("name = ?", name).Count(&count); count > 0 {
		return true
	}
	return false
}

func (m *Manager) RegisterRole(name string, description string) error {
	// 检查名称是否可用
	if m.RoleExists(name) {
		return errors.New("同名角色已经存在")
	}

	role := Role{
		Name:        name,
		Description: description,
	}

	// 添加角色
	_, err := m.AddRole(role)
	if err != nil {
		return errors.Wrap(err, "持久化角色时发生错误")
	}
	return nil
}

// AddRole 添加角色
func (m *Manager) AddRole(role Role) (*Role, error) {
	result := m.db.Create(&role)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (m *Manager) GetRole(name string) (*Role, error) {
	var role Role
	result := m.db.Where("name = ?", name).First(&role)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (m *Manager) RoleExists(name string) bool {
	var count int64
	if m.db.Model(new(Role)).Where("name = ?", name).Count(&count); count > 0 {
		return true
	}
	return false
}

func (m *Manager) DeleteRole(name string) error {
	result := m.db.Where("name = ?", name).Delete(new(Role))
	return result.Error
}
