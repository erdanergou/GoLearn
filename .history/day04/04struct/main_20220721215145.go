package main

// 结构体
import (
	"fmt"
)

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// 结构体匿名字段
// 字段比较少也比较简单的场景
// 不常用!
type dog struct {
	string
	int
}

//结构体嵌套

type animal struct {
	name string
	age  int
}

type workspace struct {
	name    string
	address string
}

type bird struct {
	an animal
}

type fish struct {
	animal // 匿名嵌套结构体
}

type mouse struct {
	animal    // 匿名嵌套结构体
	workspace //匿名嵌套结构体
}

func main() {
	// 声明一个person类型的变量p
	var zhangsan person
	// 通过字段赋值
	zhangsan.name = "张三"
	zhangsan.age = 13
	zhangsan.gender = "male"
	zhangsan.hobby = []string{"吃", "喝"}
	fmt.Println(zhangsan)
	fmt.Printf("%T\n", zhangsan)

	// 匿名结构体:多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "李四"
	s.age = 12
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	d1 := dog{
		"金毛",
		1,
	}
	fmt.Println(d1)
	fmt.Println(d1.string)
	fmt.Println(d1.int)

	b := bird{
		an: animal{
			name: "麻雀",
			age:  1,
		},
	}
	fmt.Println(b)
	fmt.Println(b.an.name) // 无法使用b.name直接访问

	f := fish{
		animal: animal{
			name: "金鱼",
			age:  2,
		},
	}
	fmt.Println(f.name) // 现在自己结构体找到这个字段,找不到就去匿名嵌套的结构体中查找

	mo := mouse{
		animal: animal{
			name:"金渐层",
			age:2,
		},
		workspace: workspace{
			name:"猫舍",
			address: "",
		},
	}
}
