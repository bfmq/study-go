package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "10.10.2.20:6379",
		Password: "tPQP4pdb6MuXcQXur4YnXgdBvzODP0lT9Yu4Y2CXXzzh0LHcmswPXf3FtPb1YOjJ",
		DB:       0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = redisdb.Ping(ctx).Result()
	return err
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("conn redis failed,err:%s\n", err)
		return
	}
	fmt.Println("reids conn!")
}
