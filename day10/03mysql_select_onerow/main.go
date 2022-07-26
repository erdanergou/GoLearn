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
	upass   string
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err%v", err)
	}
	fmt.Println("连接数据库成功")

	// 1.写查询单条记录的sql语句
	sqlStr := "SELECT uid,uname,upass,ustatus,ulevel,score FROM users where uid = ?"
	// 2.执行
	row := db.QueryRow(sqlStr, 2) //从连接池里拿出一个连接出来去数据库查询单条记录
	// row := db.QueryRow(sqlStr, 2).Scan(&u1.uid, &u1.name, &u1.upass, &u1.ustatus, &u1.ulevel, &u1.score)
	// 2.1 测试连接池
	// for i := 0; i < 10; i++ {
	// db.QueryRow(sqlStr, 1)
	// fmt.Printf("开始第%d次查询\n", i)
	// }
	// 3.拿到结果
	var u1 users
	// 必须对row对象调用scan方法，因为该方法会释放数据库连接
	row.Scan(&u1.uid, &u1.name, &u1.upass, &u1.ustatus, &u1.ulevel, &u1.score)
	fmt.Printf("%#v\n", u1)
}
