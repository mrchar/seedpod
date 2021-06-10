package db

import (
	"fmt"
	"strings"

	"github.com/mrchar/seedpod/utils"
)

// 默认设置
var defaultConfig = Config{
	Type: "postgres",
	MysqlConfig: MysqlConfig{
		Host:      "localhost",
		Port:      3306,
		User:      "seedpod",
		Password:  "seedpod-password",
		DBName:    "seedpod",
		Charset:   MysqlCharsetUTF8MB4,
		ParseTime: true,
		TimeZone:  "Local",
	},
	PostgresConfig: PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "seedpod",
		Password: "seedpod-password",
		DBName:   "seedpod",
		SSLMode:  "disable",
		TimeZone: "Asia/Shanghai",
	},
	Sqlite: SqliteConfig{
		Filepath: "sqlite.db",
	},
	Gorm: GormConfig{
		TablePrefix:   "",
		SingularTable: true,
	},
}

// Config 用于承载数据库配置
type Config struct {
	Type           string         `json:"type" yaml:"type" mapstructure:"type"`
	MysqlConfig    MysqlConfig    `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	PostgresConfig PostgresConfig `json:"postgres" yaml:"postgres" mapstructure:"postgres"`
	Sqlite         SqliteConfig   `json:"sqlite" yaml:"sqlite" mapstructure:"sqlite"`
	Gorm           GormConfig     `json:"gorm" yaml:"gorm" mapstructure:"gorm"`
}

// MysqlConfig Mysql数据库配置
type MysqlConfig struct {
	Host      string `json:"host" yaml:"host" mapstructure:"host"`
	Port      int    `json:"port" yaml:"port" mapstructure:"port"`
	User      string `json:"user" yaml:"user" mapstructure:"user"`
	Password  string `json:"password" yaml:"password" mapstructure:"password"`
	DBName    string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	Charset   string `json:"charset" yaml:"charset" mapstructure:"charset"`
	ParseTime bool   `json:"parsetime" yaml:"parsetime" mapstructure:"parsetime"`
	TimeZone  string `json:"timezone" yaml:"timezone" mapstructure:"timezone"`
}

// Validate 验证数据库配置完整性
func (c MysqlConfig) Validate() error {
	return utils.ValidateStruct(c)
}

// Datasource 返回数据库连接地址
func (c MysqlConfig) Datasource() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.Charset,
		strings.Title(fmt.Sprintf("%t", c.ParseTime)),
		c.TimeZone,
	), nil
}

// PostgresConfig PostgresSQL 数据库配置
type PostgresConfig struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DBName   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	SSLMode  string `json:"sslmode" yaml:"sslmode" mapstructure:"sslmode"`
	TimeZone string `json:"timezone" yaml:"timezone" mapstructure:"timezone"`
}

// Validate 验证数据库配置完整性
func (c PostgresConfig) Validate() error {
	return utils.ValidateStruct(c)
}

// Datasource 返回数据库连接地址
func (c PostgresConfig) Datasource() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	), nil
}

type SqliteConfig struct {
	Filepath string `json:"filepath" yaml:"filepath" mapstructure:"filepath"`
}

func (c SqliteConfig) Datasource() (string, error) {
	return c.Filepath, nil
}

type GormConfig struct {
	TablePrefix   string `json:"tablePrefix" yaml:"tablePrefix" mapstructure:"tablePrefix"`
	SingularTable bool   `json:"singularTable" yaml:"SingularTable" mapstructure:"singularTable"`
}
