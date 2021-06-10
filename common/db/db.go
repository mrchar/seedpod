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
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Gorm.TablePrefix,
			SingularTable: config.Gorm.SingularTable,
		},
	}
	var dsn string
	switch strings.ToLower(config.Type) {
	case TypePostgres:
		dsn, err = config.PostgresConfig.Datasource()
		if err != nil {
			return nil, err
		}
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	case TypeMysql:
		dsn, err = config.MysqlConfig.Datasource()
		if err != nil {
			return nil, err
		}
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	case TypeSqlite:
		dsn, err = config.Sqlite.Datasource()
		if err != nil {
			return nil, err
		}
		db, err = gorm.Open(sqlite.Open(dsn), gormConfig)
	default:
		return nil, errors.New("不支持的数据库")
	}

	return db, err
}
