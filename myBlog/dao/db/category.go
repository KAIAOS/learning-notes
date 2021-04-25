package db

import (
	"github.com/jmoiron/sqlx"
	"myBlog/model"
)

//插入单条数据
func InsertCategory(category *model.Category)(categoryId int64, err error){
	sqlstr := "insert into category(category_name,category_no) value(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil{
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

//获取单条数据
//注意返回值定义了一个category结构体的指针 代码相应变化
//DB的查询函数都应传指针进去
func GetCategory(id int64)(category *model.Category, err error){
	category = &model.Category{}
	sqlstr := "select id,category_name,category_no from category where id = ?"
	err = DB.Get(category, sqlstr, id)
	return
}

//获取多个分类
func GetCategoryList(categoryIds []int64)(categorys []*model.Category, err error){
	sqlstr, args, err := sqlx.In("select id, category_name,category_no from category where id in(?)", categoryIds)
	DB.Select(&categorys, sqlstr, args...)
	return
}

//获取所有分类
func GetAll()(categorys []*model.Category, err error){
	sqlstr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categorys, sqlstr)
	return
}