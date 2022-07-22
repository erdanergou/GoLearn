package main

type animal interface{
	move()
	eat()
}

type cat struct{
	name string
	feet int8
}

type chicken struct{
	name string
}


func main(){

}