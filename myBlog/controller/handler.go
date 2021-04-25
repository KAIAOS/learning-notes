package controller

import (
	"github.com/gin-gonic/gin"
	"myBlog/service"
)

func IndexHandle(c *gin.Context){

	list, err := service.GetAllCategoryList()
	if err != nil{
		c.JSON(500,gin.H{"message":"error","status":500})
	}
	c.JSON(200,list)

}