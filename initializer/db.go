package initializer

import (
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"thinkprinter/models"
)

func initDB() error {
	var err error
	// 连接数据库
	models.DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{
		Logger: slogGorm.New(),
	})
	if err != nil {
		return err
	}
	// 迁移 schema
	err = models.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	slog.Info("数据库初始化连接成功")
	return nil
}
