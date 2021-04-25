package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secKill/service"
)

func MiaoShaHandler(c *gin.Context){
	name := c.DefaultQuery("name","ccc")
	err := service.ReduceProductSendMsg(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "抢购成功",
		})
	}
}