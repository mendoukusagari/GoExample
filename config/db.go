package config

import (
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(GetDB),
)

func GetDB() *gorm.DB {
	dsn := "root:qwerty123@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
