package main

import (
	"github.com/gin-gonic/gin"
	"go-hub/bootstrap"
)

func main() {
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":7001")
	if err != nil {
		panic("启动失败，err: " + err.Error())
	}
}
