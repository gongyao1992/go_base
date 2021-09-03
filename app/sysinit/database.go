package sysinit

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	config "micro/upperspective/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	readBb *gorm.DB
	writeBb *gorm.DB
)

func init() {

	var err error

	// 获取读库的配置
	readC := config.Config.DB.Read

	// 读数据库
	dsnRead := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", readC.User, readC.Password, readC.Host, readC.Port, readC.Db)
	readBb, err = gorm.Open("mysql", dsnRead)
	if err != nil {
		panic(err)
	} else {
		//readBb.LogMode(true)
		//readBb.SetLogger(DbLogger)
		readBb.DB().SetMaxOpenConns(10)
	}

	writeC := config.Config.DB.Write

	// 读数据库
	dsnWrite := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", writeC.User, writeC.Password, writeC.Host, writeC.Port, writeC.Db)
	writeBb, err = gorm.Open("mysql", dsnWrite)
	if err != nil {
		panic(err)
	} else {
		writeBb.DB().SetMaxOpenConns(10)
	}

	return
}

// GetDbRead 获取读库
func GetDbRead() *gorm.DB {
	if readBb == nil {
		panic("readBb point is nil")
	}
	return readBb
}

// GetDbWrite 获取写库
func GetDbWrite() *gorm.DB {
	if writeBb == nil {
		panic("writeBb point is nil")
	}
	return writeBb
}

// InitDbLog 初始化数据库日志
func InitDbLog(d *gorm.DB, logger *log.Logger)  {
	d.LogMode(true)
	d.SetLogger(logger)
}