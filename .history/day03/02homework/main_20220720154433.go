package main

import "unicode"

func main() {
	// 1.判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s1 := "我了个乖乖omg"
	// 1.依次获取字符
	for _, c := range s1 {
		if unicode.Is(unicode.Han,c){
			
		}
	}
	// 2. 判断当前字符是否为汉字
	// 3. 出现汉字的数目求和
}
