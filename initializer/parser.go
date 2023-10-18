package initializer

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	. "thinkprinter/models"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
}

func initParser() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}

func newConfig() {
	C = Config{
		Core: struct {
			Port     int
			Bind     string
			LogLevel string
		}{
			Port:     5204,
			Bind:     "0.0.0.0",
			LogLevel: "info",
		},
		Print: struct{ WordExePath string }{
			WordExePath: "C:\\\\Program Files\\\\Microsoft Office\\\\root\\\\Office16\\\\WINWORD.EXE",
		},
		Security: struct {
			PasswordSalt  string
			JWTSecret     string
			JWTExpiration int
		}{
			PasswordSalt:  "ThinkPrinter",
			JWTSecret:     uuid.New().String(),
			JWTExpiration: 30 * 60,
		},
	}

	viper.Set("Core", C.Core)
	viper.Set("Print", C.Print)
	viper.Set("Security", C.Security)

	err := viper.SafeWriteConfig()
	if err != nil {
		slog.Error("配置文件初始化失败", "err", err)
		os.Exit(114514)
	}

	slog.Warn("配置文件初始化成功，程序即将退出，请修改配置文件后重新启动")
	os.Exit(0)
}
