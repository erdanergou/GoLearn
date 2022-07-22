package main

type animal interface {
	move()
	eat(food string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func main() {

}
