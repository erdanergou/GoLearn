package main

// 使用值接收者和指针接收者的区别？

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func main() {

}
