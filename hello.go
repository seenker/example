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
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s",server,post,database,user,password)
	if isdebug{
		fmt.Println(connString)
	}

	//建立连接
	conn, err := sql.Open("mssql",connString)
	if err !=nil{
		log.Fatal("Open Connection failed:",err.Error())
	}
	
	defer conn.Close()



}