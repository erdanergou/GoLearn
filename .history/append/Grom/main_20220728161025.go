package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户信息
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

	// // 查询
	var u = new(Person)
	db.Model(&u)
	db.First(&u, 1)
	// db.Raw("select id,name,age from person where id = 2").Scan(&u)
	fmt.Printf("%#v\n", u)

	// 更新
	// db.Model(&u).Update("age", 12)
	db.Raw("update person set age = 13 where id = 2")
	// // 删除
	// db.Delete(&u)
}
