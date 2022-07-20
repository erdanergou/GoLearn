package main


//闭包

func adder() func(int) int{
	var x int
	return func(y int) int {
		x += y
	}
}

func main(){

}