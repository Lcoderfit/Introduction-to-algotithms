package main

/*
1.缓存系统，减轻MySQL压力
2.计数器：微博抖音关注数和粉丝数
3.热门排行榜
4.实现队列

redis持久化：RDB和AOF
*/

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func initRedis() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "124541",
		DB:       0,
	})
	// 判断是否正确连接
	_, err = redisClient.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("ping redis error: %s\n", err)
		return
	}
	fmt.Println("连接成功")

	// zset
	key := "rank"
	items := []redis.Z{
		{Score: 90, Member: "PHP"},
		{Score: 95, Member: "Golang"},
		{Score: 83, Member: "Python"},
		{Score: 87, Member: "Java"},
	}
	redisClient.ZAdd(key, items...)
}
