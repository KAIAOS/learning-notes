package model

type Orders struct{
	Id int `db:"id"`
	ProductId int `db:"product_id"`
	UserName string `db:"user_name"`
}
