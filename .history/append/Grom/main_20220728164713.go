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
	db.LogMode(true)
	// 查询
	// var u Person
	// var ps []Person

	// 查询单个对象
	// 获取第一条记录（主键升序）
	// db.Table("person").First(&u) //先使用Table指明要操作的表
	// 获取一条记录，没有指定排序字段
	// db.Table("person").Take(&u)  //先使用Table指明要操作的表
	// 获取最后一条记录（主键降序）
	// db.Table("person").Last(&u)
	// result := db.Table("person").First(&u)
	// fmt.Println(result.RowsAffected)
	// db.Raw("select id,name,age from person where id = 2").Scan(&u)
	// fmt.Printf("%#v\n", u)

	// 查询多个对象
	// ret := db.Table("person").Find(&ps)
	// fmt.Printf("%v\n", ret.RowsAffected)
	// fmt.Printf("%#v\n", ps)

	// 条件查询
	// ret := db.Table("person").Where("id = ?", 1).First(&u)

	// ret := db.Table("person").Where(&Person{Name: "张三"}).First(&u)

	// ret := db.Table("person").Where(map[string]interface{}{"name": "张三", "age": 22}).Find(&u)
	// fmt.Printf("%v\n", u)
	// fmt.Printf("%v\n", ret.SubQuery())
	// 更新
	ret := db.Table("person").Model(&Person{}).Where("id = ?", 1).Update("name", "王五")
	// // 删除
	// db.Delete(&u)
}
