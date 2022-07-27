package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go连接mysql

// 定义一个全局数据库连接池
var db *sql.DB

// 构造与数据表对应的结构体

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

// 插入数据
func insert(u1 users) {
	// 1.写sql语句
	sqlStr := `insert into users(uid,uname,upass,ustatus,ulevel,score) value(?,?,?,?,?,?)`
	// 2.exec
	ret, err := db.Exec(sqlStr, u1.uid, u1.name, u1.upass, u1.ustatus, u1.ulevel, u1.score)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id: ", theId)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into users(uname,upass) value(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err: %v\n", err)
		return
	}
	//后续只需要拿到stmt去执行一些操作
	defer stmt.Close()
	var m = map[string]string{
		"a": "123",
		"b": "1234",
		"c": "12345",
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err%v", err)
	}
	fmt.Println("连接数据库成功")
	// queryOne(1)
	// queryMore(2)
	// u := users{
	// name:    "李四",
	// 	upass:   123456,
	// 	ustatus: 1,
	// 	ulevel:  0,
	// 	score:   80,
	// 	uid:     3,
	// }
	// insert(u)
	prepareInsert()
}
