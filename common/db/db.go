package db

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/schema"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var defaultDB *gorm.DB

func Default() *gorm.DB {
	var err error
	if defaultDB == nil {
		err = viper.UnmarshalKey("database", &defaultConfig)
		if err != nil {
			logrus.Panic(err)
		}

		logrus.Debugf("数据库配置: %+v", defaultConfig)

		defaultDB, err = New(defaultConfig)
		if err != nil {
			logrus.Panic(err)
		}
	}
	return defaultDB
}

func New(config Config) (db *gorm.DB, err error) {
	var dsn string
	var dialector gorm.Dialector
	switch strings.ToLower(config.Type) {
	case TypePostgres:
		dsn, err = config.PostgresConfig.Datasource()
		if err != nil {
			return nil, err
		}
		dialector = postgres.Open(dsn)
	case TypeMysql:
		dsn, err = config.MysqlConfig.Datasource()
		if err != nil {
			return nil, err
		}
		dialector = mysql.Open(dsn)
	case TypeSqlite:
		dsn, err = config.Sqlite.Datasource()
		if err != nil {
			return nil, err
		}
		dialector = sqlite.Open(dsn)

	default:
		return nil, errors.New("不支持的数据库")
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Gorm.TablePrefix,
			SingularTable: config.Gorm.SingularTable,
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
