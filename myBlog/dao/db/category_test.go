package db

import (
	"myBlog/model"
	"testing"
)

func init(){
	//parseTime=true 将mysql时间类型自动解析为go结构体中的时间类型
	dns := "root:hanKAI1998.@tcp(localhost:3306)/tx_exercise?parseTime=true"
	err := Init(dns)
	if err!=nil{
		panic(err)
	}
}

func TestInsertCategory(t *testing.T) {
	var c = model.Category{
		CategoryName: "new11",
		CategoryNo: 111,
	}
	categoryid, err := InsertCategory(&c)
	if err != nil {
		panic(err)
	}
	t.Logf("id is : %d", categoryid)
}

func TestGetCategory(t *testing.T) {
	category, err := GetCategory(1)
	if err != nil {
		panic(err)
	}

	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	categoryids := []int64{1,2,3}
	all, err := GetCategoryList(categoryids)
	if err != nil{
		panic(err)
	}
	for _,v :=range all{
		t.Logf("category:%#v", v)
	}
}

func TestGetAll(t *testing.T) {
	all, err := GetAll()
	if err != nil{
		panic(err)
	}
	for _,v := range all{
		t.Logf("is:  %#v",v)
	}
}