package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,err: ", err)
		return
	}
	//创建表，自动迁移（把结构体与数据表进行对应）
	db.AutoMigrate(&Person{})

	//查询
	var p1 person
	db.First(&p1) // 查询第一个
	fmt.Printf("%#v\n", p1)
}
