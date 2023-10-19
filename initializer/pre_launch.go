package initializer

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"thinkprinter/models"
	"thinkprinter/printer"
)

func PreLaunch() {
	// 初始化日志
	initLogger()

	// 初始化配置解析器
	err := initParser()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			newConfig()
		}
		slog.Error("配置文件解析失败", "err", err)
		slog.Warn("如果你不知道如何修复这个问题，请删除config.toml文件，然后重新启动程序。请注意，您将丢失所有配置与密钥")
		panic(err)
	}
	updateLogger()
	path := models.C.Print.WordExePath
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			slog.Error("WINWORD.EXE 文件不存在，请检查配置文件！")
		}
		panic(err)
	}

	// 初始化数据库
	err = initDB()
	if err != nil {
		slog.Error("数据库初始化连接失败", "err", err)
		panic(err)
	}

	// 启动打印机协程
	go printer.PrintWorker()
}
