package db

import (
	"github.com/liuyuexclusive/utils/appconfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //使用mysql数据库
	"github.com/sirupsen/logrus"
)

// Open 操作MYSQL数据库
func Open(f func(*gorm.DB) error) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error(r)
		}
	}()

	config := appconfig.MustGet()

	gdb, err := gorm.Open("mysql", config.ConnStr)

	defer gdb.Close()
	gdb.LogMode(false)
	gdb.SingularTable(true)
	err = f(gdb)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
