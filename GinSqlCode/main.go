package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入驱动
	"log"
)

func main() {
	connStr := "root:123456@tcp(127.0.0.1:3306)/ginsql"
	//1.链接数据库
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//创建数据库表
	//person:id,name,age
	//_, err = db.Exec("create table person (" +
	//	"id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1" +
	//	");")
	//
	//if err!=nil {
	//	log.Fatal(err.Error())
	//	return
	//}else {
	//	fmt.Println("数据库表创建成功")
	//}

	//插入数据到数据库表

	_,err=db.Exec("insert into person(name,age)" +
		"value(?,?);","Jack",20)
	if err!=nil {
		log.Fatal(err.Error())
	}else {
		fmt.Println("数据插入成功")
	}

	//查询数据库
	rows,err:=db.Query("select id,name,age from person;")
	if err!=nil {
		log.Fatal(err.Error())

	}

scan:
	if rows.Next(){
		person:=new(Person)
		err:=rows.Scan(&person.Id,&person.Name,&person.Age)
		if err!=nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id,person.Name,person.Age)
		goto scan
	}
	//按行获取数据
	//->
	//Davie 18
	//Tom 15
	//Jack 20
	//Lily 25
	//->


}

type Person struct {
	Id int
	Name string
	Age int
}