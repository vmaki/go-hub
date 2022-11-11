package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-hub/boot"
	"go-hub/config"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "", "默认加载 settings.yml 文件，若指定 --env=testing 加载的是 settings.testing.yml 文件")
	flag.Parse()
	boot.SetupConfig(env)

	boot.SetupLogger()

	r := gin.New()

	boot.SetupRoute(r)

	err := r.Run(":" + cast.ToString(config.Cfg.Application.Port))
	if err != nil {
		panic("启动失败，err: " + err.Error())
	}
}
