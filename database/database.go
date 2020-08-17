package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

var db *sqlx.DB

func Open(auth, pwd, addr, dbname string, port int) error {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		auth, pwd, addr, port, dbname)
	db_, err := sqlx.Open("mysql", connstr)
	if err != nil {
		return err
	}
	db = db_



	return nil


}

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return db.Get(dest, query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}


func Queryx(query string,args ...interface{})(*sqlx.Rows, error) {
	return db.Queryx(query,args...)
}

type DataTable []map[string]interface{} //列表通用返回值
type DataMap map[string]interface{} //单条通用返回值
// 根据SQL查询列表
func QueryListBySql(sql string)(DataTable,error){
	rows,err := Queryx(sql)
	if err != nil{
		return nil,err
	}
	columns,_ := rows.Columns()
	columnsLength := len(columns)
	cache := make([]interface{},columnsLength)//临时存储每行数据
	for index,_ := range cache{//为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}//返回的切片
	for rows.Next(){
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i,data := range cache{
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		FmtDataUnit8ToString(item)//必须将unit8 转为 字符串
		list = append(list,item)
	}
	return list,nil
}
// 根据SQL查询单行
func QueryBySql(sql string)(DataMap,error){
	list,err := QueryListBySql(sql)
	if err != nil {
		return nil,nil
	}
	if len(list)>0 {
		return list[0],nil
	}
	return nil,nil
}
//结果转换 []unit8 转 字符串
func FmtDataUnit8ToString(r map[string]interface{}) {
	for k, v := range r {
		switch v.(type) {
		case []uint8:
			arr := v.([]uint8)
			r[k] = string(arr)
		case nil:
			r[k] = ""
		}
	}
}
