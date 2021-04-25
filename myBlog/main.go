package main

import (
	"github.com/gin-gonic/gin"
	"myBlog/controller"
	"myBlog/dao/db"
)

func main(){
	r := gin.Default()
	dns := "root:hanKAI1998.@tcp(localhost:3306)/tx_exercise?parseTime=true"
	err := db.Init(dns)
	if err!=nil{
		panic(err)
	}

	r.GET("/", controller.IndexHandle)

	r.Run(":8000")
}
