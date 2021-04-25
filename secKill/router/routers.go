package router

import (
	"github.com/gin-gonic/gin"
	"secKill/controller"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/miaosha", controller.MiaoShaHandler)
	return r
}