package main

//递归

func jiecheng(n int)(ret int){
	if(n > 0){
		ret = n * jiecheng(n - 1)
	}
	return ret
}

func main() {
	jiecheng(3)
}
