package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*Go操作mysql
database/sql
原生支持连接池，是并发安全的
这个标准库没有具体实现，只是列出了一些需要第三方库实现的具体内容
*/

// go连接mysql

// 定义一个全局对象
// var db *sql.DB

func main() {
	// dsn 数据库名
	dsn := "root:123456@tcp(localhost:3306)/xiaomissm"
	// 不会校验账号密码是否正确
	//
	db, err := sql.Open("mysql", dsn)
	if err != nil { //dsn格式不正确时会报错
		fmt.Printf("dsn %s invalid,err:%v ", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Open %s failed,err:%v ", dsn, err)
		return
	}
	fmt.Println("连接成功")
}
