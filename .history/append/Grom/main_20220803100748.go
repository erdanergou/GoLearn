package main

import (
	"gorm_now/mytable"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:root@(localhost:3306)/manage_archives?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	// db.AutoMigrate(&HonoraryArchives{})
	// 查询
	// var u mytable.Person

	var cs []mytable.Company
	// var ps []mytable.Person
	// var c mytable.Company

	// 使用TableName方法
	// db.Model(&u).Where("id = ?", 2).First(&u)

	// db.Select("name").Where("id = ?", 2).First(&u)

	// 预加载

	// db.First(&u, 2)
	// 一对一模式
	// db.Model(&mytable.Person{}).Preload("Job").First(&u, 5)
	// fmt.Println(u)

	// 一对多，但无法打印Person的Job
	// db.Model(&mytable.Company{}).Preload("Persons").Find(&cs)
	// fmt.Println(cs)

	// 链式预加载(先预加载Persons.Job)
	// db.Model(&mytable.Company{}).Preload("Persons.Job").Preload("Persons").Find(&cs)
	// fmt.Println(cs)

	//多对多
	// db.Model(&mytable.Company{}).Preload("Jobs").Find(&cs)
	// fmt.Println(cs)


	// join连接
	db.Model(&mytable.Company).Joins("Jobs")




	// 没有定义tablename方法
	// 查询单个对象
	// 获取第一条记录（主键升序）
	// db.Table("person").First(&u) //先使用Table指明要操作的表
	// db = db.Table("person").Where("age = ?", 12)
	// db.Table("person").Where("id = ?", 2).Find(&u)
	// ps = append(ps, Person{Name: "王五", Age: 18, Id: 2})
	// ps = append(ps, Person{Name: "赵六", Age: 19, Id: 4})

	// 获取一条记录，没有指定排序字段
	// db.Table("person").Take(&u)  //先使用Table指明要操作的表
	// 获取最后一条记录（主键降序）
	// db.Table("person").Last(&u)
	// result := db.Table("person").First(&u)
	// fmt.Println(result.RowsAffected)
	// db.Raw("select id,name,age from person where id = 2").Scan(&u)
	// db = db.Table("person").Model(&u)
	// where := make(map[string]interface{})
	// where["age"] = 18
	// db.Where(where).Find(&ps)
	// fmt.Printf("%#v\n", u)
	// fmt.Printf("%#v\n", ps)

	// 查询多个对象
	// db.Table("person").Where("id = ?", 2).Find(&ps)

	// fmt.Printf("%v\n", ret.RowsAffected)
	// fmt.Printf("%#v\n", ps)
	// fmt.Printf("%v\n", len(ps))

	// result := map[string]interface{}{}
	// db.Model(&Person{}).First(&result)
	// 条件查询
	// ret := db.Table("person").Where("id = ?", 1).First(&u)

	// ret := db.Table("person").Where(&Person{Name: "张三"}).First(&u)

	// ret := db.Table("person").Where(map[string]interface{}{"name": "张三", "age": 22}).Find(&u)
	// fmt.Printf("%v\n", u)
	// fmt.Printf("%v\n", ret.SubQuery())
	// 更新
	// db.Table("person").Model(&Person{}).Where("id = ?", 2).Update("name", "李四")
	// db.Table("person").Model(&Person{}).Update("name", "张三")
	// // 删除
	// db.Delete(&u)
	// db.Table("person").Transaction(func(tx *gorm.DB) error {
	// 	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	// 	if err := tx.Create(&Person{Name: "Giraffe"}).Error; err != nil {
	// 		// 返回任何错误都会回滚事务
	// 		return err
	// 	}

	// 	if err := tx.Create(&Person{Name: "Lion"}).Error; err != nil {
	// 		return err
	// 	}

	// 	// 返回 nil 提交事务
	// 	return nil
	// })
	// where := make(map[string]interface{}, 10)
	// where["id"] = 4
	// where["age"] = 12
	// db.Table("person").Where(where["id"]).Delete(&u)
}
