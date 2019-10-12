package dbutil

import (
	"utils/configutil"

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

	config := configutil.MustGet()

	db, err := gorm.Open("mysql", config.ConnStr)

	defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)
	err = f(db)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
