package config

import (
	"github.com/spf13/viper"
	"go-hub/config"
	"os"
)

var v *viper.Viper

func init() {
	// 1. 初始化 viper 库
	v = viper.New()

	// 2. 相关配置
	v.SetConfigType("yml")   // 配置类型
	v.AddConfigPath(".")     // 环境变量配置文件查找的路径，相对于 main.go
	v.SetEnvPrefix("appEnv") // 设置环境变量前缀，用以区分 Go 的系统环境变量
	v.AutomaticEnv()         // 读取环境变量（支持 flags）
}

func LoadEnv(env string) {
	envPath := "settings.yml"
	if env != "" {
		envPath = "settings." + env + "+.yml"
	}

	// 检查配置文件是否存在
	if _, err := os.Stat(envPath); err != nil {
		panic("启动失败，err: 配置文件 " + envPath + " 不存在")
	}

	// 加载配置文件
	v.SetConfigName(envPath)
	if err := v.ReadInConfig(); err != nil {
		panic("启动失败，err: 加载配置文件失败，" + err.Error())
	}

	err := v.Unmarshal(&config.Cfg)
	if err != nil {
		panic("启动失败，err: 读取配置文件失败，" + err.Error())
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}
