package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	// 配置文件名为当前工作目录下的 config.toml
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件失败", err)
		log.Println("需要初始化配置文件!")

		// 初始化配置文件
		err = initConfig()
		if err != nil {
			log.Fatalln("初始化配置文件失败", err)
		}

		log.Println("初始化配置文件成功，请重新启动程序")
		os.Exit(0)
	}

	// 将配置文件中的配置项映射到结构体中
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Println("映射配置文件失败", err)
		log.Println("需要初始化配置文件!")

		// 初始化配置文件
		err = initConfig()
		if err != nil {
			log.Fatalln("初始化配置文件失败", err)
		}

		log.Println("初始化配置文件成功，请重新启动程序")
		os.Exit(0)
	}
}

func initConfig() error {
	// 设置默认配置
	viper.Set("core", Core{
		Port: 8080,
		Bind: "0.0.0.0",
	})
	viper.Set("security", Security{
		PasswordSalt:  "thinkPrinter",
		JWTSecret:     "thinkPrinter",
		JWTExpiration: 3600,
	})

	// 写入配置文件
	log.Println("正在生成配置文件...")
	err := viper.WriteConfigAs("config.toml")
	if err != nil {
		return err
	}

	return nil
}
