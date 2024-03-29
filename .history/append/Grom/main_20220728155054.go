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

	// 自动迁移
	db.AutoMigrate(&User{})

	// u1 := UserInfo{1, "七米", "男", "篮球"}
	// u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// // 查询
	var u = new(User)
	db.First(u)
	fmt.Printf("%#v\n", u)

	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)

	// // 更新
	// db.Model(&u).Update("hobby", "双色球")
	// // 删除
	// db.Delete(&u)
}
