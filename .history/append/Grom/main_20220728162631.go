package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Person 用户信息
type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Table(&Person)
	// 查询
	var u Person
	ret := db.First(&u, 1)
	// esult := db.Where("id = ?", "1").First(&u)
	// db.Raw("select id,name,age from person where id = 2").Scan(&u)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%v\n", ret)

	// 更新
	// db.Model(&u).Update("age", 12)
	db.Raw("update person set age = 13 where id = 2")
	// // 删除
	// db.Delete(&u)
}
