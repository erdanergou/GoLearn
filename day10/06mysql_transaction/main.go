package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type users struct {
	uid     int
	name    string
	upass   int
	ustatus int
	ulevel  int
	score   int
}

func initDB() (err error) {
	// dsn 数据库名
	dsn := "root:123456@tcp(localhost:3306)/xiaomissm"
	// 不会校验账号密码是否正确
	//
	db, err = sql.Open("mysql", dsn)
	if err != nil { //dsn格式不正确时会报错
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

func transactionDemo() {
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err %v\n", err)
		return
	}
	// 执行多个sql操作
	sqlStr1 := `update users set  score=score + 2 where uid = 1`
	sqlStr2 := `update users set  score=score - 2 where uid = 2`
	// sql1执行
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行sql1出错,要回滚")
		return
	}
	// sql2执行
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行sql2出错,要回滚")
		return
	}

	// 上述操作都成功,就提交
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("提交出错了,要回滚")
		return
	}
	fmt.Println("事务成功了")

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err%v", err)
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}
