package main

import (
	"fmt"
	"strings"
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
		if unicode.Is(unicode.Han, c) {
			// 3. 出现汉字的数目求和
			num++
		}
	}
	fmt.Println(num)

	// 2. how do you do单词出现的次数
	s2 := "how do you do"
	// 把字符串按空格切割
	s3 := strings.Split(s2, " ")
	// 遍历切片存储到一个map.
	m1 = make(map[string]int, 10)
	for _, w := range s3 {
		// 如果当前map不存在该key则初始化值为1
		// 如果map中存在这个key，则加1
		if count,ok := 
		fmt.Println(w)
	}

	// 累加出现次数
}
