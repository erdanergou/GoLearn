package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int map[string]int) {
	m := make(map[string]int, 10)
	for _, v := range users {
		m[v] = 0
	}

	for _, v := range users {
		for _, w := range v {
			switch w {
			case 'e':
				m[v] += 1
				coins -= 1
			case 'E':
				m[v] += 1
				coins -= 1
			case 'i':
				m[v] += 2
				coins -= 2
			case 'I':
				m[v] += 2
				coins -= 2
			case 'o':
				m[v] += 3
				coins -= 3
			case 'O':
				m[v] += 3
				coins -= 3
			case 'u':
				m[v] += 4
				coins -= 4
			case 'U':
				m[v] += 4
				coins -= 4
			}

		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
