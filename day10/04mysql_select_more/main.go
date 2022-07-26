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

func queryOne(id int) {
	// 1.写查询单条记录的sql语句
	sqlStr := "SELECT uid,uname,upass,ustatus,ulevel,score FROM users where uid = ?"
	// 2.执行
	row := db.QueryRow(sqlStr, id) //从连接池里拿出一个连接出来去数据库查询单条记录
	// row := db.QueryRow(sqlStr, 2).Scan(&u1.uid, &u1.name, &u1.upass, &u1.ustatus, &u1.ulevel, &u1.score)
	// 3.拿到结果
	var u1 users
	// 必须对row对象调用scan方法，因为该方法会释放数据库连接
	row.Scan(&u1.uid, &u1.name, &u1.upass, &u1.ustatus, &u1.ulevel, &u1.score)
	fmt.Printf("%#v\n", u1)
}

func queryMore(id int) {
	// 1.sql语句
	sqlStr := "SELECT uid,uname,upass,ustatus,ulevel,score FROM `users`"
	var rows *sql.Rows
	var err error
	if id < 0 {
		rows, err = db.Query(sqlStr)
	} else {
		sqlStr = sqlStr + "where uid >= ?"
		rows, err = db.Query(sqlStr, id)
	}
	// 2.执行
	if err != nil {
		fmt.Printf("exec %s query failed,err:%v\n", sqlStr, err)
		return
	}
	// 3.一定要关闭rows
	defer rows.Close()
	// 4.循环取值
	for rows.Next() {
		var u1 users
		err := rows.Scan(&u1.uid, &u1.name, &u1.upass, &u1.ustatus, &u1.ulevel, &u1.score)
		if err != nil {
			fmt.Printf("scan DB failed,err:%v\n", err)
		}
		fmt.Printf("%#v\n", u1)
	}

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

// 修改数据
func update(u1 users) {
	sqlStr := `update users set  uname=?,upass=?,ustatus=?,ulevel=?,score=? where uid = ?`
	ret, err := db.Exec(sqlStr, u1.name, u1.upass, u1.ustatus, u1.ulevel, u1.score, u1.uid)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id  failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据", n)
}

// 删除
func delete(id int) {
	sqlStr := `delete from users where uid = ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id  failed,err:%v\n", err)
		return
	}
	fmt.Printf("影响了%d条数据", n)
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
	// update(u)
	delete(4)
}
