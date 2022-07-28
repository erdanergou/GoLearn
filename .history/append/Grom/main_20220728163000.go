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
	// 查询
	var u Person
	db.Table("person").Take(&u) //先使用Table指明要操作的表
	db.Table("person").Take(&u) //先使用Table指明要操作的表
	// db.Raw("select id,name,age from person where id = 2").Scan(&u)
	fmt.Printf("%#v\n", u)

	// 更新
	// db.Table("person").Model(&u).Update("age", 22)
	// // 删除
	// db.Delete(&u)
}
