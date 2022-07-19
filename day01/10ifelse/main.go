package main

import "fmt"

// if条件判断
func main() {
	// age := 19
	// if age < 18 {
	// 	fmt.Println("未成年不得入内")
	// } else if age > 18 && age < 30 {
	// 	fmt.Println("冲一百送五十")
	// } else {
	// 	fmt.Println("年纪大的别上网了")
	// }


	//age的作用域在if条件判断语句中生效
	if age := 19; age>18{
		fmt.Println("澳门首家线上赌场开业了")
	}else{
		fmt.Println("好好学习天天向上")
	}

	
}
