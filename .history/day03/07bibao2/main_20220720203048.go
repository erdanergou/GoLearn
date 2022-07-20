package main


//闭包

func adder() func(int) int{
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main(){
	ret := adder()
	ret2 := ret(200)
	
}