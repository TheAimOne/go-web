package connection

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RDS *redis.Client

func Init() {
	RDS = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	err := RDS.Set(ctx, "key", "value111", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RDS.Get(ctx, "key").Result()

	fmt.Println("the dinga here+++", val)
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return RDS
}
