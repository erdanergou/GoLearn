package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s1 := "我了个乖乖omg"
	num := 0
	// 1.依次获取字符
	for _, c := range s1 {
		// 2. 判断当前字符是否为汉字
		if unicode.Is(unicode.Han,c){
			// 3. 出现汉字的数目求和
			num ++
		}
	}
	fmt.Println(num)

	// 2. how do you do单词出现的次数
	s2 := "how do you do"
	// 把字符串按空格切割
	s3 := s
	// 遍历切片存储到一个map
	// 累加出现次数
}
