package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func InitMysql() (err error){
	dns := "root:hanKAI1998.@tcp(localhost:3306)/tx_exercise?parseTime=true"
	DB, err = sqlx.Open("mysql",dns)
	if err!=nil{
		return err
	}

	err = DB.Ping()
	if err!= nil{
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return
}

func Close() {
	DB.Close()
	Redisdb.Close()
	MQCH.Close()
	conn.Close()
}