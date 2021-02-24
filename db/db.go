package db

import (
	"github.com/liuyuexclusive/utils/appconfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //使用mysql数据库
	"github.com/sirupsen/logrus"
)

type DB struct {
	*gorm.DB
}

// Open 操作MYSQL数据库
func Open(f func(*DB) error) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error(r)
		}
	}()

	config := appconfig.MustGet()

	gdb, err := gorm.Open("postgres", config.ConnStr)

	defer gdb.Close()
	gdb.LogMode(true)
	gdb.SingularTable(true)
	err = f(&DB{gdb})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
