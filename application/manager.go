package application

import (
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/db"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var defaultManager *Manager

type Manager struct {
	db             *gorm.DB
	accountManager *account.Manager
}

// DefaultManager 使用默认的数据库创建Manager
func DefaultManager() *Manager {
	if defaultManager == nil {
		defaultManager = NewManager(db.Default(), account.DefaultManager())
	}
	return defaultManager
}

// NewManager 创建Manager
func NewManager(db *gorm.DB, accountManager *account.Manager) *Manager {
	return &Manager{
		db:             db,
		accountManager: accountManager,
	}
}

// RegisterApplication 注册应用程序
func (m *Manager) RegisterApplication(name string, description ...string) (string, error) {
	app := Application{Name: name}
	if len(description) > 0 {
		app.Description = description[0]
	}

	new, err := m.AddApplication(app)
	if err != nil {
		err = errors.Wrap(err, "持久化应用时发生错误")
		return "", err
	}

	return new.AppId, err
}

// AddApplication 添加应用程序
func (m *Manager) AddApplication(application Application) (*Application, error) {
	result := m.db.Create(&application)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &application, nil
}

// GetApplicationByAppId 根据appId获取应用程序
func (m *Manager) GetApplicationByAppId(appId string) (*Application, error) {
	var application Application
	result := m.db.Where("app_id = ?", appId).First(&application)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &application, nil
}

// DeleteApplication 删除appId对应的应用程序
func (m *Manager) DeleteApplication(appId string) error {
	result := m.db.Where("appId=?", appId).Delete(new(Application))
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetAccountByAppIdAndAccountName(appId, accountName string) (*Account, error) {
	app, err := m.GetApplicationByAppId(appId)
	if err != nil {
		return nil, errors.Wrap(err, "检查Application失败")
	}

	act, err := m.accountManager.GetAccount(accountName)
	if err != nil {
		return nil, errors.Wrap(err, "检查Account失败")
	}

	var appAct Account
	result := m.db.Where("application_uuid=? AND account_uuid = ?", app.UUID, act.UUID).First(&appAct)
	if err := result.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "检查ApplicationAccount失败")
	}

	if err == gorm.ErrRecordNotFound {
		logrus.Debug("ApplicationAccount不存在，将创建")
		appAct = Account{ApplicationUUID: app.UUID, AccountUUID: act.UUID}
		result := m.db.Create(appAct)
		if err := result.Error; err != nil {
			return nil, errors.Wrap(err, "创建ApplicationAccount失败")
		}
	}

	return &appAct, nil
}
