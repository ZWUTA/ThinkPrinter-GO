package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"thinkPrinter/entity"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	// 连接数据库
	DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// 迁移 schema
	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}

	log.Println("数据库连接成功")
	return nil
}
