package main

import "fmt"

//方法

type dog struct {
	name string
}

//构造函数
func newDag(name string) dog {
	return dog{
		name: name,
	}
}

func wang() {
	fmt.Println("汪汪汪")
}

func main() {

}
