package main

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // "_" 引入后面的包名 而不直接使用里面的定义的函数、变量、资源等

)

func main() {
	db, err := sql.Open("mysql", "root:ly@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil || db == nil {
		fmt.Println("连接错误")
		return
	}
	//name := "df"
	listsql := sq.Select("name, id").
		From("user").
		Where(fmt.Sprintf("name like '%s'", "%adf% and (SELECT GROUP_CONCAT(schema_name) FROM information_schema." +
			"schemata)"))
		//Where(sq.Like{"name": name})
	sqlStr, _, _ := listsql.ToSql()
	fmt.Println(sqlStr)
	//fmt.Printf("%v\n", args)

	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var id int32
		if err := rows.Scan(&name, &id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s , id: %v \n", name, id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
