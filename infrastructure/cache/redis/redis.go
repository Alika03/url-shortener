package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"shortener/config"
	"strconv"
	"sync"
)

var (
	rDb    *redis.Client
	onceDb sync.Once
)

func GetDbConnection() *redis.Client {
	onceDb.Do(func() {
		dbNumber, _ := strconv.Atoi(config.GetConfig().CacheDbParams.Db)
		rDb = redis.NewClient(&redis.Options{
			Addr:     config.GetConfig().CacheDbParams.Host + ":" + config.GetConfig().CacheDbParams.Port,
			Password: config.GetConfig().CacheDbParams.Password,
			DB:       dbNumber,
		})

		_, err := rDb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln("ping error:", err)
		}
		//if pong != "" {
		//	log.Fatalln("empty pong")
		//}
	})
	return rDb
}
