package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type person struct {
	id   int
	name string
	age  int
}

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:root@localhost:3306/test?charset=utf8mb4&parseTime=True&loc=local")
	if err != nil {
		fmt.Println("open mysql failed,err: ", err)
		return
	}
	//创建表，自动迁移（把结构体与
}
