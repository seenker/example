package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/denisenkom/go-mssqldb"

)


func main(){
	var isdebug = true
	var server = "localhost"
	var port = 1433
	var user = "sa"
	var password = "1qaz2wsx"
	var database = "kaoqin"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s",server,port,database,user,password)
	if isdebug{
		fmt.Println(connString)
	}

	//建立连接
	conn, err := sql.Open("mssql",connString)
	if err !=nil{
		log.Fatal("Open Connection failed:",err.Error())
	}
	
	defer conn.Close()

	//产生查询语句的Statement
	stmt, err := conn.Prepare("select [USERID] ,[BADGENUMBER],[SSN],[NAME],[GENDER] from [USERINFO]")
	if err !=nil{
		log.Fatal("Prepare failed:",err.Error())
	}
	defer stmt.Close()

	//通过Statement执行语句
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:",err.Error())
	}
	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{},len(cols))
	for i :=0;i < len(cols); i++ {
		colsdata[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	//建立一个列数组
	for rows.Next(){
		rows.Scan(colsdata...) //将查到的数据写入到这行
		PrintRow(colsdata)   //
	}
	defer rows.Close()
}

//
func PrintRow(colsdata []interface{}){
	for _, val := range colsdata{
		switch v:=(*(val.(*interface{}))).(type){
		case nil:
			fmt.Print("NULL")
		case bool:
			if v{
				fmt.Print("true")
			}else{
				fmt.Print("false")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2017-01-30 15:15:15.999"))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}