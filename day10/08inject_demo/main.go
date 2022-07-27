package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// SQL注入

var db *sqlx.DB

type user struct {
	Name string
	Age  int
	Id   int
}

func initDB() (err error) {
	// dsn 数据库名
	dsn := "root:123456@tcp(localhost:3306)/test"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn %s invalid,err:%v ", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Open %s failed,err:%v ", dsn, err)
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// sql注入实例
func sqlInjectDemo(name string) {
	// 自己拼接sql语句字符串
	sqlStr := fmt.Sprintf("select name,age,id from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err: %v\n", err)
		return
	}
	// SQL 注入的几种实例
	// sqlInjectDemo("zs")
	// sqlInjectDemo("x' or '1 = 1 #")
	sqlInjectDemo("x' union select * from user #")
}
