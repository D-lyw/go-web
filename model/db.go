package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _gormDB *gorm.DB

func init() {
	//postgresDB, err := sql.Open("postgres", "postgres://go_web_user:iW5gPaUFB9gCqsMcOwhJvjIrAnbvERkH@dpg-cmpnicmn7f5s73dfoqog-a.singapore-postgres.render.com/go_web")
	//if err != nil {
	//	panic(err)
	//}
	dsn := "postgres://go_web_user:iW5gPaUFB9gCqsMcOwhJvjIrAnbvERkH@dpg-cmpnicmn7f5s73dfoqog-a.singapore-postgres.render.com/go_web"

	_gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = _gormDB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("Auto Migrate Schema Error", err.Error())
	}

	//_gormDB.
	sqlDB, _ := _gormDB.DB()

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
}

func GetDB() *gorm.DB {
	return _gormDB
}
