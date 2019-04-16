package model

import (
	"bishe/spider/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/core/errors"
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
	connectMySQL(config.MySQL)
}

func connectMySQL(config config.DBConfig) {
	username, password, protocol, address, DBname := config.Username, config.Password, config.Protocol, config.Address, config.DBname
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", username, password, protocol, address, DBname))
	if err != nil {
		panic(err)
	}
}
// 插入
func Insert(data PositionInfo) error {
	if data.IsMiss {
		return errors.New("missing data item")
	}
	isUnique, err := isUniquePid(data.Pid)
	if err != nil {
		return err
	}
	if !isUnique {
		return errors.New("repeat data")
	}

	sql := `INSERT INTO position_info(pid, position, salary_low, salary_high, city, workyear, level, education, company_name, industry, finance_stage, company_size, detail, location)
						VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Pid, data.Positon,
				data.Salary.Low,
				data.Salary.High,
				data.City, data.WorkYear,
				data.Level, data.Education,
				data.CompanyName, data.Industry,
				data.FinanceStage, data.CompanySize,
				data.Detail, data.Location)
	if err != nil {
		return err
	}
	return nil
}
// 查找
func isUniquePid(pid string) (bool, error) {
	sql := `SELECT pid FROM position_info where pid=?`
	row, err := db.Query(sql, pid)
	if err != nil {
		return false, err
	}
	cn := 0
	for row.Next() {
		cn++;
	}
	if cn > 0 {
		return false, nil
	}
	return true, nil
}
// 删除
func Delete() {

}
// 更新
func Update() {

}