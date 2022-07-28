package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
	Id   int    `gorm:"column:id"`
}

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,err: ", err)
		return
	}
	defer db.Close()
	fmt.Printf("%#v\n", db)
	//查询
	var p1 Person
	db.First(&p1) // 查询第一个
	fmt.Printf("%#v\n", p1)

}
