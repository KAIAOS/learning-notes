package service

import (
	"myBlog/dao/db"
	"myBlog/model"
)

func GetAllCategoryList()(categoryList []*model.Category, err error){
	categoryList, err = db.GetAll()
	if err != nil{
		return
	}
	return
}