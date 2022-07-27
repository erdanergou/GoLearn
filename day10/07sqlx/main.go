package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// sqlx
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	var u1 user
	sqlStr := `select id,name,age from user where id = 1;`
	// 通过反射来对结构体进行赋值，所以需要将字段定义为公共方法
	db.Get(&u1, sqlStr)
	fmt.Printf("%#v\n", u1)

	var userList = make([]user, 10)
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,%v\n", err)
	}
	fmt.Printf("userlist %#v\n", userList)
}
