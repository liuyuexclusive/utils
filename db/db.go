package db

import (
	"github.com/yuexclusive/utils/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Open() (*gorm.DB, error) {
	cfg := config.MustGet()

	gdb, err := gorm.Open("postgres", cfg.ConnStr)
	if err != nil {
		return nil, err
	}

	gdb.LogMode(true)
	gdb.SingularTable(true)
	return gdb, nil
}
