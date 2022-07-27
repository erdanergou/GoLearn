package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// redis
/*

 */

// 声明全局redis对象
var redisdb *redis.Client

func initRedis() (err error) {
	// go-redis 库中使用 redis.NewClient 函数连接 Redis 服务器
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功")

	// zset 有序集合
	key := "rank"
	items := []redis.Z{
		{Score: 97, Member: "PHP"},
		{Score: 98, Member: "Python"},
		{Score: 99, Member: "Golang"},
		{Score: 100, Member: "Java"},
	}

	// 把元素都追加到key
	redisdb.ZAdd(key, (items)...)

	// 把java的score减2
	newScroe, err := redisdb.ZIncrBy(key, -2, "Java").Result()
	if err != nil {
		fmt.Printf("z zincrby failed,err:%v\n", err)
		return
	}
	fmt.Println(newScroe)
}
