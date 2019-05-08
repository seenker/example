package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)
type AccessRegion struct{
	region_id			int64
	provider_id 		int64
	region_name			string
	sub_region_name		string
	billing_region_name	string
	description			string
}
func main(){
	var isdebug = true
	var server = "localhost"
	var port = 1433
	var user = "sa"
	var password = "1qaz2wsx"
	var database = "kaoqin"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s",server,port,database,user,password)
	
	//建立连接
	db,err := sql.Open("mssql",connString)
	if err !=nil{
		log.Fatal("Open Connection failed:",err.Error())
	}
	defer db.Close()

	//通过连接对象执行查询
	rows, err :=db.Query(`select [USERID] ,[BADGENUMBER],[SSN],[NAME],[GENDER] from [USERINFO]`)
	if err !=nil{
		log.Fatal("Query failed:",err.Error())
	}
	defer rows.Close()

	//通过连接对象执行查询
	colNames, _ :=rows.Columns()
	var cols = make([]interface{},len(colNames))

	

}