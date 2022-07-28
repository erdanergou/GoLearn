package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Name string
	Age  int
	Id   int
}

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,err: ", err)
		return
	}
	defer db.Close()
	//查询
	var u1 User
	db.First(&u1) // 查询第一个
	fmt.Printf("%#v\n", u1)

}
