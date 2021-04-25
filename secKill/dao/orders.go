package dao

import "secKill/model"

//插入一条订单数据

func InsertOrders(o *model.Orders)(id int64,err error){
	sql := "insert into orders(product_id,user_name) value(?,?)"
	res, err := DB.Exec(sql, o.ProductId, o.UserName)
	if err != nil{
		return
	}
	id, _ = res.LastInsertId()
	return
}

func GetOrdersById(productId int)(orders []*model.Orders, err error){
	orders = []*model.Orders{}
	sql := "select id,product_id,user_name from orders where product_id=?"
	err = DB.Select(&orders, sql, productId)
	return
}