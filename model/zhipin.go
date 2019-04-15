package model

import (
	"bishe/spider/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type PositionInfo struct {
	Pid string
	Positon string
	Salary config.Salary
	City string
	WorkYear string
	Level int
	Education string
	CompanyName string
	Industry string
	FinanceStage string
	CompanySize string
	Detail string
	Location string
	IsMiss bool
}
// 数据存储

var (
	db *sql.DB; err error
)

func init(){
	connect(config.MySQL)
}

func connect(config config.DBConfig) {
	dbname, username, password, protocol, address, tablename := config.DBname, config.Username, config.Password, config.Protocol, config.Address, config.Tablename
	db, err = sql.Open(dbname, fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", username, password, protocol, address, tablename))
	if err != nil {
		panic(err)
	}
}
// 插入
func Insert(data PositionInfo) {

}
// 查找
func Select() {

}
// 删除
func Delete() {

}
// 更新
func Update() {

}