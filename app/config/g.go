package config

import (
	"github.com/gongyao1992/go-util/helper"
	"gopkg.in/ini.v1"
)

const configFile = "conf.ini"

// mysql 的配置结构
type mysqlConfig struct {
	Host string
	User string
	Password string
	Db string
	Port string
}

var Config = struct {
	DB struct{
		Read mysqlConfig
		Write mysqlConfig
	}

	Log struct{
		Dir string `default:"/var/logs/vip/"`
		DbDir string `default:"/var/logs/db/"`
	}

	FileDir string `default:""`
}{}

func init() {
	// 加载配置文件
	configPath := helper.GetConfigFile(configFile)
	f, err := ini.Load(configPath)
	if err != nil {
	}

	var c *mysqlConfig
	c, err = getMySqlConfig(f, "mysql-read")
	if err != nil {
	}
	if c == nil {
		panic("")
	}
	Config.DB.Read = *c

	c, err = getMySqlConfig(f, "mysql-write")
	if err != nil {
	}
	if c == nil {
		panic("")
	}
	Config.DB.Write = *c

	err = initLogDir(f)
	if err != nil {
		panic(err.Error())
	}

	err = initFileDir(f)
	if err != nil {
		panic(err.Error())
	}
}

func initLogDir(f *ini.File) error {

	section, err := f.GetSection("log")
	if err != nil {
		return err
	}
	key, err := section.GetKey("dir")
	if err != nil {
		return err
	}
	Config.Log.Dir = key.String()

	dbkey, err := section.GetKey("dbdir")
	if err != nil {
		return err
	}
	Config.Log.DbDir = dbkey.String()

	return nil
}

func getMySqlConfig(f *ini.File, name string) (*mysqlConfig, error) {
	section, err := f.GetSection(name)
	if err != nil {
		return nil, err
	}

	return &mysqlConfig{
		Host:     section.Key("host").String(),
		User:     section.Key("user").String(),
		Password: section.Key("password").String(),
		Db:       section.Key("db").String(),
		Port:     section.Key("port").String(),
	}, nil
}

func initFileDir(f *ini.File) error {
	section, err := f.GetSection("paths")
	if err != nil {
		return err
	}
	key, err := section.GetKey("tempfile")
	if err != nil {
		return err
	}
	Config.FileDir = key.String()

	return nil
}